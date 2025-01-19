package db

import (
	"context"
	"fmt"
	"log"
	"project_green/config"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

// InitDB initializes the database connection using pgx
func InitDB() {
	dbConfig, err := config.LoadDBConfig()
	if err != nil {
		log.Fatalf("Failed to load database configuration: %v", err)
	}

	// Construct the connection string
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	// Connect to the database
	DB, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	// Test the connection
	if err := DB.Ping(context.Background()); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Successfully connected to the database!")
}
