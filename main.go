package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/marina-yamaguti/go-gorm-crud/config"
	"github.com/marina-yamaguti/go-gorm-crud/handlers"
)

func main() {
	// Initialize Database
	config.Connect()

	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Create a Gin router
	router := gin.Default()

	// Initialize product handler
	productHandler := &handlers.ProductHandler{DB: config.DB}

	// Setup routes
	router.POST("/products", productHandler.CreateProduct)
	router.GET("/products", productHandler.GetProducts)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)

	// Start the server
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
