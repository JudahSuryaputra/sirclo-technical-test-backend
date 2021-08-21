package main

import (
	"sirclo-technical-test-backend/http"
	"sirclo-technical-test-backend/models/db"
)

func main() {
	sqlDB := db.Conn()

	http.RunServer(sqlDB)
}