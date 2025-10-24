package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	ServerAddr  string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	dbUser := os.Getenv("POSTGRES_USER")
	if dbUser == "" {
		return nil, fmt.Errorf("POSTGRES_USER is not set")
	}

	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	if dbPassword == "" {
		return nil, fmt.Errorf("POSTGRES_PASSWORD is not set")
	}

	dbName := os.Getenv("POSTGRES_DB")
	if dbName == "" {
		return nil, fmt.Errorf("POSTGRES_DB is not set")
	}

	serverAddr := os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		return nil, fmt.Errorf("SERVER_ADDR is not set")
	}

	databaseURL := fmt.Sprintf("postgres://%s:%s@db:5432/%s?sslmode=disable", dbUser, dbPassword, dbName)

	return &Config{
		DatabaseURL: databaseURL,
		ServerAddr:  serverAddr,
	}, nil
}
