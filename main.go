package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"project_green/db"
	"project_green/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}

	// Initialize the database
	db.InitDB()
	defer db.DB.Close(context.Background())

	// Initialize chi router
	r := chi.NewRouter()

	// Register routes
	handlers.DeviceRoutes(r)
	handlers.SensorDataRoutes(r)

	// Get the port from the environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
