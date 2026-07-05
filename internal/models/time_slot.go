package models

import "gorm.io/gorm"

type TimeSlot struct {
	gorm.Model
	Label     string `gorm:"size:50;not null"` // e.g. "10:00 AM - 11:00 AM"
	StartTime string `gorm:"size:10;not null"` // HH:MM
	EndTime   string `gorm:"size:10;not null"`
	IsActive  bool   `gorm:"default:true"`
}
