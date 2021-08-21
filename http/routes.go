package http

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"sirclo-technical-test-backend/http/handler"

	"github.com/gorilla/mux"
)

func RunServer(sqlDB *sql.DB) {
	port := "5000"

	r := mux.NewRouter()

	insertWeight := handler.InsertWeight{DB: sqlDB}
	getWeights := handler.GetWeights{DB: sqlDB}
	getWeight := handler.GetWeight{DB: sqlDB}
	updateWeight := handler.UpdateWeight{DB: sqlDB}
	r.Handle("/weight", insertWeight).Methods(http.MethodPost)
	r.Handle("/weights", getWeights).Methods(http.MethodGet)
	r.Handle("/weight/{date}", getWeight).Methods(http.MethodGet)
	r.Handle("/weight/{id}", updateWeight).Methods(http.MethodPut)
	
	fmt.Printf("\nServer starting on Port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}