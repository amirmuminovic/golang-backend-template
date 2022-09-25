package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func connectToDb(ac AppConfig) *sql.DB {
	db, err := sql.Open(ac.driverType, ac.connectionString)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
