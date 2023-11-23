package handlers

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"project-management-application/db"
	"project-management-application/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

func downloadAndCompressImage(imageURL, destinationDir string) (string, error) {
	// Download the image
	response, err := http.Get(imageURL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Get the file name from the URL
	fileName := getFileNameFromURL(imageURL)

	// Create the file in the specified destination directory
	filePath := filepath.Join(destinationDir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Copy the content of the image to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	// Compress the image
	err = compressImage(filePath, filepath.Join(destinationDir, fileName+"_compressed.jpg"), 300, 300)
	if err != nil {
		return "", err
	}

	compressedImagePath := filepath.Join(destinationDir, fileName+"_compressed.jpg")
	fmt.Printf("Image downloaded and compressed successfully: %s\n", compressedImagePath)
	return compressedImagePath, nil
}

func compressImage(inputPath, outputPath string, maxWidth, maxHeight uint) error {
	// Open the input image file
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// Resize the image
	resized := resize.Resize(maxWidth, maxHeight, img, resize.Lanczos3)

	// Create the output file
	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Encode the resized image to the output file
	err = jpeg.Encode(out, resized, nil)
	if err != nil {
		return err
	}

	return nil
}

// Function to extract the file name from a URL
func getFileNameFromURL(url string) string {
	lastIndex := strings.LastIndex(url, "/")
	if lastIndex != -1 {
		return url[lastIndex+1:]
	}
	return url
}



// CreateProduct handler with image download and compression for a single image

// func CreateProduct(c *gin.Context) {
// 	var productInput models.Product
// 	if err := c.ShouldBindJSON(&productInput); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}


// 	var user models.User
// 	if err := db.DB.First(&user, productInput.UserID).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id or user not found"})
// 		return
// 	}

// 	// Download and compress product image
// 	destinationPath:="C:/Users/sushi/OneDrive/Desktop/Zocket-Assignment-with-go-programming/project_management_application/images/";
// 	compressedImagePath, err := downloadAndCompressImage(productInput.ProductImages,destinationPath)
// 	fmt.Println("compressedfile is ",compressedImagePath)
// 	fmt.Println("image", productInput.ProductImages)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to download and compress product image"})
// 		return
// 	}

// 	// Update productInput with compressed image path
// 	productInput.CompressedProductImages = compressedImagePath

// 	// Create a new product in the database
// 	if err := db.DB.Create(&productInput).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, productInput)
// }




func CreateProduct(c *gin.Context){
	var productInput models.Product
	if err := c.ShouldBindJSON(&productInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.DB.First(&user, productInput.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id or user not found"})
		return
	}

	if err := db.DB.Create(&productInput).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, productInput)
}

func GetProduct(c *gin.Context) {
	productID := c.Param("id")

	var product models.Product
	if err := db.DB.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// func ListProducts(c *gin.Context) {
// 	var products []models.Product

// 	// Retrieve all products from the database (replace this with your database logic)
// 	// For simplicity, we'll assume you have a global 'db' variable from db.go
// 	if err := db.DB.Find(&products).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
// 		return
// 	}

// 	// Fetch associated user details for each product
// 	for i := range products {
// 		var user models.User
// 		if err := db.DB.First(&user, products[i].UserID).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user details for a product"})
// 			return
// 		}
// 		// Do not expose sensitive user details, adjust this based on your needs
// 		products[i].User = user
// 	}

// 	c.JSON(http.StatusOK, products)
// }

func UpdateProduct(c *gin.Context) {
	productID := c.Param("id")

	var existingProduct models.Product
	if err := db.DB.First(&existingProduct, productID).Error; err != nil {
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
	if err := db.DB.Model(&existingProduct).Updates(updatedProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, existingProduct)
}

func DeleteProduct(c *gin.Context) {
	productID := c.Param("id")

	var existingProduct models.Product
	if err := db.DB.First(&existingProduct, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Delete the product from the database (replace this with your database logic)
	// For simplicity, we'll assume you have a global 'db' variable from db.go
	if err := db.DB.Delete(&existingProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
