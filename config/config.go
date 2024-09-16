// config/config.go
package config

import (
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// Config struct to hold configuration values
type Config struct {
	DBURL               string
	Port                string
	JWTAuthSecret       string
	JWTAuthExpInHour    time.Duration
	JWTRefreshSecret    string
	JWTRefreshExpInHour time.Duration
}

// LoadConfig loads configuration from .env file and returns the Config object and database connection
func LoadConfig() (*Config, *sqlx.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, nil, err
	}

	dbURL := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	jwtAuthExpInHour, err := strconv.Atoi(os.Getenv("JWT_AUTH_EXP_IN_HOUR"))
	if err != nil {
		return nil, nil, err
	}

	jwtRefreshExpInHour, err := strconv.Atoi(os.Getenv("JWT_REFRESH_EXP_IN_HOUR"))
	if err != nil {
		return nil, nil, err
	}

	cfg := &Config{
		DBURL:               os.Getenv("DATABASE_URL"),
		Port:                os.Getenv("PORT"),
		JWTAuthSecret:       os.Getenv("JWT_AUTH_SECRET"),
		JWTRefreshSecret:    os.Getenv("JWT_REFRESH_SECRET"),
		JWTAuthExpInHour:    time.Duration(jwtAuthExpInHour) * time.Hour,
		JWTRefreshExpInHour: time.Duration(jwtRefreshExpInHour) * time.Hour,
	}

	if err := cfg.Validate(); err != nil {
		return nil, nil, err
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
