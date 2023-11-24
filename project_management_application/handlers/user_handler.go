package handlers

import (
	"net/http"
	"project-management-application/db"
	"project-management-application/models"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"message":"Hello World!"})
}


func CreateUserHandler(c *gin.Context){
	var userInput models.User

	// Bind JSON data to User struct
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&userInput).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Respond with the created user
	c.JSON(http.StatusOK, userInput)
}