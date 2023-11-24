package handlers

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"project-management-application/db"
	"project-management-application/models"
	"time"

	"github.com/gin-gonic/gin"
)

// sampleData := `{
// 	"user_id": "1",
// 	"product_name": "Sample Product 2",
// 	"product_description": "Description of the sample product 2",
// 	"product_images": "https://upload.wikimedia.org/wikipedia/commons/thumb/3/3a/Cat03.jpg/1025px-Cat03.jpg",
// 	"product_price": 29.99
// }`


func CreateProductHandler(c *gin.Context) {

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

	fmt.Println("user", user)
	fmt.Println("product", productInput)
	fmt.Println("product images", productInput.ProductImages)

	if err := db.DB.Create(&productInput).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	// Download the image

	go func() {
		res, err := downloadAndCompressImage(productInput.ProductImages)
        if err != nil {
            // Handle the error (e.g., log it)
            fmt.Printf("Error downloading image: %s\n", err.Error())
            return
        }

        fmt.Printf("Image downloaded and compressed: %s\n", res)

        // Update the productInput with the path to the downloaded image
        productInput.CompressedProductImages = res

        // Update the product in the database with the image path
        if err := db.DB.Save(&productInput).Error; err != nil {
            // Handle the error (e.g., log it)
            fmt.Printf("Error updating product with image path: %s\n", err.Error())
            return
        }

        // Optionally, you can log a message indicating that the image processing is complete
        fmt.Println("Image processing completed for product:", productInput.ID)
	}()

	c.JSON(http.StatusOK, productInput)
}

func downloadAndCompressImage(imageURL string) (string, error) {

	fmt.Printf("Downloading image from: %s\n", imageURL)

	// Generate a unique filename
	fileName := generateUniqueFileName()

	// Define the destination path for the downloaded and compressed image
	destinationPath := "C:/Users/sushi/OneDrive/Desktop/Zocket-Assignment-with-go-programming/project_management_application/images/" + fileName

	// Download the image
	response, err := http.Get(imageURL)
	if err != nil {
		fmt.Printf("Error downloading image: %s\n", err.Error())
		return "", err
	}
	defer response.Body.Close()

	// Create the file
	file, err := os.Create(destinationPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Compress the image (replace this with your compression logic)
	// For simplicity, we'll just copy the content for demonstration purposes
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	fmt.Printf("Image downloaded and compressed successfully: %s\n", destinationPath)
	return destinationPath, nil
}

// Function to generate a unique filename using timestamp and random string
func generateUniqueFileName() string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	randomString := generateRandomString(8) // You can adjust the length of the random string
	return fmt.Sprintf("%d_%s.jpg", timestamp, randomString)
}

// Function to generate a random string
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
