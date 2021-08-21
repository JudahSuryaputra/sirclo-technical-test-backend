package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"sirclo-technical-test-backend/models/requests"
	"strconv"

	"github.com/gorilla/mux"
)

type UpdateWeight struct {
	DB *sql.DB
}

func (c UpdateWeight) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathVariable := mux.Vars(r)
	weightID, _ := strconv.Atoi(pathVariable["id"])

	var request requests.UpdateWeightRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	request.Difference = request.Max - request.Min

	sqlStatement := `UPDATE weights SET max=$1, min=$2, difference=$3, date=$4 WHERE id=$5`

	result, err := c.DB.Exec(sqlStatement, request.Max, request.Min, request.Difference, request.Date, weightID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
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