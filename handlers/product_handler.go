package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marina-yamaguti/go-gorm-crud/models"
	"gorm.io/gorm"
)

type ProductHandler struct {
	DB *gorm.DB
}

// CREATE
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product

	// 1. Bind JSON input to product struct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Create product in the database
	result := h.DB.Create(&product)
	
	// 3. Handle potential errors
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// 4. Return the created product
	c.JSON(http.StatusCreated, product)
}

// READ
func (h *ProductHandler) GetProducts(c *gin.Context) {
	var products []models.Product
	
	// 1. Retrieve all products from the database
	result := h.DB.Find(&products)

	// 2. Handle potential errors
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// 3. Return the list of products
	c.JSON(http.StatusOK, products)
}

// TO DO: Implement UPDATE and DELETE handlers