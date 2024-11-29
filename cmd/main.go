package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iyiola-dev/numeris/internal/db"
	"github.com/iyiola-dev/numeris/internal/models"
	"github.com/iyiola-dev/numeris/internal/routes"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Warning: .env file not found")
	}
	// Set Gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Connect to database
	db.ConnectDatabase()

	// AutoMigrate models
	err := db.DB.AutoMigrate(
		&models.User{},
		&models.Customer{},
		&models.Invoice{},
		&models.InvoiceItem{},
		&models.ActivityLog{},
		&models.PaymentDetails{},
	)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Migrations completed successfully!")

	// Initialize router
	router := routes.SetupRouter()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
