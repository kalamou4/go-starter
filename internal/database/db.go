package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
	"time"
)

func Connect(databaseURL string) (*sql.DB, error) {
	// Open database connection
	db, err := sql.Open("sqlite3", databaseURL)
	if err != nil {
		return nil, err
	}

	// Verify connection with a ping
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Set connection pool settings (optional but recommended)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
