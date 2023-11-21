package handlers

import (
	"net/http"
	"project-management-application/db"
	"project-management-application/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/go-playground/validator/v10"
)


func Hello(c *gin.Context){
	c.JSON(http.StatusOK,"Hello world!");
}


// CreateUser creates a new user in the database.
func CreateUser(c *gin.Context) {
	var userInput models.User

	// Validate the input data
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perform stricter input validation using a validation library
	validate := validator.New()
	err := validate.Struct(userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the user in the database
	if err := db.DB.Create(&userInput).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Respond with the created user
	c.JSON(http.StatusOK, userInput)
}

// GetUser retrieves a user from the database by ID.
func GetUser(c *gin.Context) {
	userID := c.Param("id")

	// Convert ID to uint
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		if err == db.DB.Error {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	c.JSON(http.StatusOK,user);
}