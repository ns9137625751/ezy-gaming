package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"uniqueIndex;size:150;not null"`
	Password string `gorm:"size:255;not null"`
	Phone    string `gorm:"size:20"`
}
