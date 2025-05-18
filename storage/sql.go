package storage

import (
	"database/sql"
	"fmt"
	"log"
)

func Connect() *sql.DB {
	connStr := "user=nwt_user password=networth dbname=networth host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	fmt.Println("Connected to the database successfully.")

	return db
}
