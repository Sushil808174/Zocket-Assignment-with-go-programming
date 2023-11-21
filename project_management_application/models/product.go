package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	UserID                  string   `json:"user_id" gorm:"not null"`
	ProductName             string   `json:"product_name" gorm:"not null"`
	ProductDescription      string   `json:"product_description"`
	ProductImages           []string `json:"product_images" gorm:"type:json"`
	ProductPrice            float64  `json:"product_price"`
	CompressedProductImages []string `json:"compressed_product_images" gorm:"type:json"`
	User                    User     `json:"user"`
}