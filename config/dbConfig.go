package config

import (
	"log"
	"os"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

// LoadDBConfig loads the database configuration based on the environment
func LoadDBConfig() (*DBConfig, error) {

	var config DBConfig

	config.User = os.Getenv("DB_USER")
	config.Password = os.Getenv("DB_PASSWORD")
	config.Host = os.Getenv("DB_HOST")
	config.Port = os.Getenv("DB_PORT")
	config.Name = os.Getenv("DB_NAME")

	log.Printf("DBConfig: Host=%s, Port=%s, Name=%s", config.Host, config.Port, config.Name)

	return &config, nil
}
