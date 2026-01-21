// @title           Product API
// @version         1.0
// @description     This is a simple CRUD API for managing products.
// @host            localhost:8080
// @BasePath        /

package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/marina-yamaguti/go-gorm-crud/config"
	_ "github.com/marina-yamaguti/go-gorm-crud/docs"
	"github.com/marina-yamaguti/go-gorm-crud/handlers"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	router.PUT("/products/:id", productHandler.UpdateProduct)
	router.DELETE("/products/:id", productHandler.DeleteProduct)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	log.Printf("Swagger UI available at http://localhost:%s/swagger/index.html", port)

	// Start the server
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
