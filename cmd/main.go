// cmd/main.go
package main

import (
	"crowdfunding/config"
	"crowdfunding/internal/v1/auth"
	"crowdfunding/pkg/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, db, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Database connection established")
	r := gin.Default()

	// Middleware
	r.Use(middleware.CORSMiddleware())

	// API versioning
	authService := auth.NewService(db)
	authHandler := auth.NewHandler(authService)

	v1 := r.Group("/api/v1")
	{
		auth.RegisterRoutes(v1, authHandler)
	}

	// Start server
	log.Printf("Starting server on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
