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

// CreateProduct godoc
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

// UpdateProduct godoc
// @Summary      Update an existing product
// @Description  Update a product by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id       path      int             true  "Product ID"
// @Param        product  body      models.Product  true  "Product Data"
// @Success      200      {object}  models.Product
// @Failure      404      {object}  map[string]string
// @Router       /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	// 1. Check if product exists
	if err := h.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// 2. Bind new data
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. Use Updates to only change the fields provided in the JSON
	h.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Delete a product by its ID
// @Tags         products
// @Param        id   path      int  true  "Product ID"
// @Success      204  {object}  nil
// @Failure      404  {object}  map[string]string
// @Router       /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	// 1. Delete the product
	result := h.DB.Delete(&models.Product{}, id)

	// 2. Check if any row was affected
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// 3. Return no content status
	c.Status(http.StatusNoContent)
}
