package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sirclo-technical-test-backend/models/db"
	"sirclo-technical-test-backend/models/responses"
)

type GetWeights struct {
	DB *sql.DB
}

func (c GetWeights) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var weights []db.Weight
	sqlStatement := `SELECT * FROM weights ORDER BY date DESC`

	rows, err := c.DB.Query(sqlStatement)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	var maxWeight, minWeight, weightDifference, totalData int
	
	for rows.Next() {
		var weight db.Weight
		
		err := rows.Scan(&weight.ID, &weight.Max, &weight.Min, &weight.Difference, &weight.Date)

		if err != nil {
            w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
        }
		
		maxWeight = maxWeight + weight.Max
		minWeight = minWeight + weight.Min
		weightDifference= weightDifference + weight.Difference
		totalData++

		weights =  append(weights, weight)
	}

	average := responses.Average{
		Max: 		float64(maxWeight)/float64(totalData),
		Min: 		float64(minWeight)/float64(totalData),
		Difference: float64(weightDifference)/float64(totalData),
	}

	response := responses.WeightIndexResponse{
		Weights: weights,
		Average: average,
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}