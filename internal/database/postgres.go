package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Connect establishes a connection to the PostgreSQL database.
func Connect(connStr string) (*gorm.DB, error) {
	var err error
	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		return nil, err
	}

	fmt.Println("Successfully connected to the database")
	return db, err
}

// GetDB returns the database connection.
func GetDB() *gorm.DB {
	return db
}

func Close() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Error getting sql.DB from gorm.DB: %v", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}
}
