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

// / CreateProduct godoc
// @Summary      Create a new product
// @Description  Create a new product with the input payload
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      models.Product  true  "Product Data"
// @Success      201      {object}  models.Product
// @Failure      400      {object}  map[string]string
// @Router       /products [post]
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

// GetProducts godoc
// @Summary      Get all products
// @Description  Retrieve a list of all products
// @Tags         products
// @Produce      json
// @Success      200      {array}   models.Product
// @Failure      500      {object}  map[string]string
// @Router       /products [get]
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

// UPDATE
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	// TO DO: Implement UPDATE handler
}

// DELETE
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	// TO DO: Implement DELETE handler
}
