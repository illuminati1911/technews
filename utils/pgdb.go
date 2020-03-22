package utils

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	// lib/pq will
	_ "github.com/lib/pq"
)

// OpenDBConnection configures and opens database
// connection to PostgreSQL database
func OpenDBConnection() (*sql.DB, error) {
	var db *sql.DB
	var err error
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	connectionRetries, err := strconv.Atoi(os.Getenv("POSTGRES_RETRIES"))
	if err != nil {
		return nil, err
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		port,
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	for connectionRetries > 0 {
		db, err = sql.Open("postgres", psqlInfo)
		err = db.Ping()
		if err != nil {
			// logger -> SQL database could not be connected
			connectionRetries--
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	return db, err
}
