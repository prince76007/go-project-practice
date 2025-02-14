package main

import (
    "log"
    "net/http"

    "go-project/internal/config"
    "go-project/internal/database"
    "go-project/internal/routes"
)

func main() {
    // Load configuration
    cfg := config.LoadConfig()

    // Initialize database connection
    db, err := database.Connect(cfg.DatabaseURL)
    if err != nil {
        log.Fatalf("Could not connect to database: %v", err)
    }
    defer db.Close()

    // Set up HTTP routes
    router := routes.SetupRouter()

    // Start the server
    log.Printf("Starting server on port %s", cfg.ServerPort)
    if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
        log.Fatalf("Could not start server: %v", err)
    }
}