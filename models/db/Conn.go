package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "sirclo-dev"
)

func Conn() (*sql.DB) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("Cannot initialize connection to database: %v", err)
	}

	log.Printf(dsn)

	return sqlDB
}