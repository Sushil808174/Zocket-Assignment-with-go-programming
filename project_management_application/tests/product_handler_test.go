package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project-management-application/handlers"
	"project-management-application/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	// Create a sample product JSON
	productJSON := []byte(`{
		"user_id": "1",
		"product_name": "Sample Product",
		"product_description": "Sample Description",
		"product_images": "image1.jpg",
		"product_price": 19.99
	}`)

	// Create a new HTTP request with the product JSON as the request body
	req, err := http.NewRequest("POST", "/create-product", bytes.NewBuffer(productJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	recorder := httptest.NewRecorder()

	// Create a new Gin context from the response recorder
	context, _ := gin.CreateTestContext(recorder)
	context.Request = req

	// Call the CreateProduct function
	handlers.CreateProduct(context)

	// Check the HTTP status code
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Parse the response body
	var response models.Product
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// Assert the response values as needed
	assert.Equal(t, "Sample Product", response.ProductName)
	assert.Equal(t, "Sample Description", response.ProductDescription)
	// Add more assertions based on your implementation
}
