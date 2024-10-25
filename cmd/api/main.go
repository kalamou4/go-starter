package main

import (
	"log"
	"net/http"
	"time"

	"go-starter/internal/auth"
	"go-starter/internal/database"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	cfg := config.Load()
	db := database.Connect(cfg)
	defer db.Close()

	// Initialize router
	router := mux.NewRouter()

	// Initialize repositories and services
	userRepo := user.NewRepository(db)
	authService := auth.NewService(userRepo)

	// Setup routes
	auth.RegisterRoutes(router, authService)
	user.RegisterRoutes(router, userRepo)

	// Configure server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server
	log.Printf("Server starting on port %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
