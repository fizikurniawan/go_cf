// config/config.go
package config

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// Config struct to hold configuration values
type Config struct {
	DBURL string
	Port  string
}

// LoadConfig loads configuration from .env file and returns the Config object and database connection
func LoadConfig() (*Config, *sqlx.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, nil, err
	}

	dbURL := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")
	if dbURL == "" || port == "" {
		return nil, nil, os.ErrNotExist
	}

	db, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		return nil, nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, nil, err
	}

	config := &Config{
		DBURL: dbURL,
		Port:  port,
	}

	return config, db, nil
}
