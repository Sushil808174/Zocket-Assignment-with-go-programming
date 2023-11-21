package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"project-management-application/models"
)

var db *gorm.DB

// ConnectDB connects to the database and performs migrations.
func ConnectDB() {
	// Define the database connection string (DSN)
	dsn := "root:Sushil8081@tcp(127.0.0.1:3306)/product_management?parseTime=true"

	// Open a connection to the MySQL database
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	// Perform database migrations
	db.AutoMigrate(models.User{}, models.Product{})
}