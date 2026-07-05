package models

import (
	"time"

	"gorm.io/gorm"
)

type ContactMessage struct {
	gorm.Model
	Name      string     `gorm:"size:100;not null"`
	Email     string     `gorm:"size:150;not null"`
	Phone     string     `gorm:"size:20"`
	Subject   string     `gorm:"size:200"`
	Message   string     `gorm:"type:text;not null"`
	IsRead    bool       `gorm:"default:false"`
	RepliedAt *time.Time
}
