package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sirclo-technical-test-backend/models/db"

	"github.com/gorilla/mux"
)

type GetWeight struct {
	DB *sql.DB
}

func (c GetWeight) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathVariable := mux.Vars(r)
	requestedDate := pathVariable["date"]

	sqlStatement := `SELECT * FROM weights WHERE date=$1`

	row := c.DB.QueryRow(sqlStatement, requestedDate)

	var weight db.Weight
	err := row.Scan(&weight.ID, &weight.Max, &weight.Min, &weight.Difference, &weight.Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weight)
}