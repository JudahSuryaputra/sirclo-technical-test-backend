package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"sirclo-technical-test-backend/models/requests"
)

type InsertWeight struct {
	DB *sql.DB
}

func (c InsertWeight) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request requests.InsertWeightRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	weightDifference := request.Max - request.Min

	sqlStatement := `INSERT INTO weights (date, max, min, difference) VALUES ($1, $2, $3, $4)`

	result, err := c.DB.Exec(sqlStatement, request.Date, request.Max, request.Min, weightDifference)
	if err 	!= nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err 	!= nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	if rowsAffected != 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.New("row affected should be 1"))
	}

	return
}