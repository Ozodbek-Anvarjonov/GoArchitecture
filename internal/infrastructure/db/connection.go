package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect(connectionString string) error {
	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("Failed to open db connection: %w", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("Failed to ping db connection: %w", err)
	}

	log.Println("Successfully connected to db")
	return nil
}
func Close() error {
	if DB != nil {
		return DB.Close()
	}

	return nil
}
