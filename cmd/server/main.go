package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres" // or sqlite depending on your DB
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() {
	dsn := os.Getenv("DB_DSN") // e.g., "host=localhost user=postgres password=postgres dbname=mydb port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected successfully.")
}

func main() {
	// Initialize DB
	initDB()

	// Initialize router
	r := gin.Default()

	// Register routes
	RegisterRoutes(r, DB)

	// Start server
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
