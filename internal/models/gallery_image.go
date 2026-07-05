package models

import "gorm.io/gorm"

type GalleryImage struct {
	gorm.Model
	Title      string `gorm:"size:200"`
	ImageURL   string `gorm:"size:255;not null"`
	Category   string `gorm:"size:50"` // setup | events | tournaments | cafe
	IsFeatured bool   `gorm:"default:false"`
	SortOrder  int    `gorm:"default:0"`
}
