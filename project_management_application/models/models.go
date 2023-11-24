package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


type Product struct {
	gorm.Model
	UserID                  string   `json:"user_id" gorm:"not null"`
	ProductName             string   `json:"product_name" gorm:"not null"`
	ProductDescription      string   `json:"product_description"`
	ProductImages           string `json:"product_images" gorm:"type:text"`
	ProductPrice            float64  `json:"product_price"`
	CompressedProductImages string `json:"compressed_product_images" gorm:"type:text"`
  }
  
  