package repositories

import (
	"log"
	"testing"

	"go-project-practice/internal/database"

	"go-project-practice/internal/config"

	"gorm.io/gorm"
)

var testDB *gorm.DB

func SetupTestDB() {
	var err error
	cfg := config.LoadConfig()

	testDB, err = database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to test database:", err)
	}
}

func TestMain(m *testing.M) {
	SetupTestDB()
	m.Run()
}
