package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"project-management-application/handlers"
	"project-management-application/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func CreateUserTest(t *testing.T){
	req, err := http.NewRequest("POST", "/create-user", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	recorder := httptest.NewRecorder()

	// Create a new Gin context from the response recorder
	context, _ := gin.CreateTestContext(recorder)
	context.Request = req

	// Call the CreateUser function
	handlers.CreateUser(context)

	// Check the HTTP status code
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Parse the response body
	var response models.User
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}
	 fmt.Println("response ",response)
	// Assert the response values as needed
	assert.Equal(t, "John Doe", response.Name)
	assert.Equal(t, "1234567890", response.Mobile)
	// Add more assertions based on your implementation
}