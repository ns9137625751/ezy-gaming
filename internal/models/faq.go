package models

import "gorm.io/gorm"

type FAQ struct {
	gorm.Model
	Question  string `gorm:"type:text;not null"`
	Answer    string `gorm:"type:text;not null"`
	Category  string `gorm:"size:50"`
	SortOrder int    `gorm:"default:0"`
	IsActive  bool   `gorm:"default:true"`
}
