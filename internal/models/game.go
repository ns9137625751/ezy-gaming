package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	Title      string `gorm:"size:200;not null"`
	Genre      string `gorm:"size:100"`
	CoverURL   string `gorm:"size:255"`
	Platform   string `gorm:"size:20;not null"` // pc | ps5 | both
	IsFeatured bool   `gorm:"default:false"`
}
