package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name" gorm:"not null"`
	Mobile    string `json:"mobile" gorm:"unique;not null"`
	Latitude  float64
	Longitude float64
}
