package main

import (
	"log"

	"go-project-practice/internal/config"
	"go-project-practice/internal/database"
	"go-project-practice/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database connection
	_, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer database.Close()

	// Set up HTTP routes
	router := gin.Default()
	routes.RegisterQuestionRoutes(router)
	routes.RegisterProjectRoutes(router)

	// Start the server
	log.Printf("Starting server on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
