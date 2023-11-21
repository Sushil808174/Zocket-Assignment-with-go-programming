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

	// Validate that the provided user_id exists in the Users table
	userID := c.PostForm("user_id")
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	// Associate the user_id with the product
	productInput.UserID = userID

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

func ListProducts(c *gin.Context){
	var products []models.Product

	// Retrieve all products from the database (replace this with your database logic)
	// For simplicity, we'll assume you have a global 'db' variable from db.go
	if err := db.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	// Fetch associated user details for each product
	for i := range products {
		var user models.User
		if err := db.First(&user, products[i].UserID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user details for a product"})
			return
		}
		// Do not expose sensitive user details, adjust this based on your needs
		products[i].User = user
	}

	c.JSON(http.StatusOK, products)
}

func UpdateProduct(c *gin.Context) {
	productID := c.Param("id")

	var existingProduct models.Product
	if err := db.First(&existingProduct, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var updatedProduct models.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update product details in the database (replace this with your database logic)
	// For simplicity, we'll assume you have a global 'db' variable from db.go
	if err := db.Model(&existingProduct).Updates(updatedProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, existingProduct)
}


func DeleteProduct(c *gin.Context) {
	productID := c.Param("id")

	var existingProduct models.Product
	if err := db.First(&existingProduct, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Delete the product from the database (replace this with your database logic)
	// For simplicity, we'll assume you have a global 'db' variable from db.go
	if err := db.Delete(&existingProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}