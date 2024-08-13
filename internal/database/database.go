package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func InitDB(databaseURL string) (*sql.DB, error) {
	var err error
	db, err = sql.Open("postgres", databaseURL)
	if err != nil {
		log.Printf("Failed to connect to: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}

	log.Printf("Connected to database")
	return db, nil
}

func GetDB() *sql.DB {
	return db
}
