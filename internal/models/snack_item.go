package models

import "gorm.io/gorm"

type SnackItem struct {
	gorm.Model
	Name        string  `gorm:"size:150;not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"not null"`
	Category    string  `gorm:"size:50"` // beverages | snacks | meals
	ImageURL    string  `gorm:"size:255"`
	IsAvailable bool    `gorm:"default:true"`
}
