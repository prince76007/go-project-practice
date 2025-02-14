package database

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

// Connect establishes a connection to the PostgreSQL database.
func Connect(connStr string) {
    var err error
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }

    if err = db.Ping(); err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

    fmt.Println("Successfully connected to the database")
}

// Close closes the database connection.
func Close() {
    if err := db.Close(); err != nil {
        log.Fatalf("Error closing the database: %v", err)
    }
    fmt.Println("Database connection closed")
}

// GetDB returns the database connection.
func GetDB() *sql.DB {
    return db
}