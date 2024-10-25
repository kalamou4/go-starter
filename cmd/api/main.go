package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"go-starter/internal/auth"
	"go-starter/internal/config"
	"go-starter/internal/database"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	err = config.RunMigrations(cfg.MigrationURL, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize database
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize router (using standard net/http)
	mux := http.NewServeMux()

	// Initialize auth service and handler
	authService := auth.NewService(db)
	authHandler := auth.NewHandler(authService)

	// Register routes
	mux.HandleFunc("/api/auth/register", authHandler.Register)
	mux.HandleFunc("/api/auth/login", authHandler.Login)
	mux.Handle("/api/protected", auth.RequireAuth(http.HandlerFunc(authHandler.Protected)))

	// Start server
	log.Printf("Server starting on %s", cfg.ServerAddr)
	if err := http.ListenAndServe(cfg.ServerAddr, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
