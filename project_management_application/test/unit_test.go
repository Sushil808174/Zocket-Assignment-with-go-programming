package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project-management-application/db"
	"project-management-application/handlers"
	"project-management-application/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateProductHandler(t *testing.T) {
	// Initialize the Gin router
	router := gin.Default()

	// Connect to the database
	db.ConnectDB()

	// Add the route to the router
	router.POST("/create-product", handlers.CreateProductHandler)

	// Create a test user
	user := models.User{
		Name: "Test User",
		// Add other necessary fields
	}
	db.DB.Create(&user)

	// Create a test product input
	productInput := models.Product{
		UserID:             "1",
		ProductName:       "Sample Product 2",
		ProductDescription: "Description of the sample product 2",
		ProductImages:     "https://upload.wikimedia.org/wikipedia/commons/thumb/3/3a/Cat03.jpg/1025px-Cat03.jpg",
		ProductPrice:      29.99,
	}

	// Convert product input to JSON
	jsonInput, err := json.Marshal(productInput)
	if err != nil {
		t.Fatal(err)
	}

	// Create a request to simulate the API call
	req, err := http.NewRequest("POST", "/create-product", bytes.NewBuffer(jsonInput))
	if err != nil {
		t.Fatal(err)
	}

	// Set the Content-Type header to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response
	recorder := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(recorder, req)

	// Check if the status code is 200 (OK)
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Optionally, parse the response body to verify the result
	var responseProduct models.Product
	err = json.Unmarshal(recorder.Body.Bytes(), &responseProduct)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, productInput.ProductName, responseProduct.ProductName)
	assert.Equal(t, productInput.ProductDescription, responseProduct.ProductDescription)
	assert.Equal(t, productInput.ProductImages, responseProduct.ProductImages)
	assert.Equal(t, productInput.ProductPrice, responseProduct.ProductPrice)
}