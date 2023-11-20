package handlers

import (
	"net/http"
	"project-management-application/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userInput models.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new user in the database (replace this with your database logic)
	// For simplicity, we'll assume you have a global 'db' variable from db.go
	if err := db.Create(&userInput).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, userInput)
}

func GetUser(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}