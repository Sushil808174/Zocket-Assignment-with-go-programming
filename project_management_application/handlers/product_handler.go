package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var productInput models.Product
	if err := c.ShouldBindJSON(&productInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new product in the database (replace this with your database logic)
	// For simplicity, we'll assume you have a global 'db' variable from db.go
	if err := db.Create(&productInput).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, productInput)
}

func GetProduct(c *gin.Context) {
	productID := c.Param("id")

	var product models.Product
	if err := db.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}