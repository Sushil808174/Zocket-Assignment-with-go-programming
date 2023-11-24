package db

import (
	"fmt"
	"project-management-application/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:Sushil8081@tcp(127.0.0.1:3306)/product_management?parseTime=true"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	// Perform database migrations
	DB.AutoMigrate(models.User{}, models.Product{})
}