package models

import "gorm.io/gorm"

type PS5Console struct {
	gorm.Model
	StationNumber   int     `gorm:"uniqueIndex;not null"`
	Name            string  `gorm:"size:100"`
	Tier            string  `gorm:"size:20"` // Premium | Standard
	TVSize          string  `gorm:"size:50"`
	ControllerCount int     `gorm:"default:2"`
	HourlyRate      float64 `gorm:"not null"`
	Status          string  `gorm:"size:20;default:available"`
	ImageURL        string  `gorm:"size:255"`
	IsActive        bool    `gorm:"default:true"`
}
