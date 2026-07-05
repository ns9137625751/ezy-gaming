package models

import "gorm.io/gorm"

type AdminUser struct {
	gorm.Model
	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"size:150;uniqueIndex;not null"`
	Password string `gorm:"size:255;not null"`
	Role     string `gorm:"size:20;default:admin"` // admin | gamer
	IsActive bool   `gorm:"default:true"`
}
