package models

import "gorm.io/gorm"

type GamingPC struct {
	gorm.Model
	StationNumber int     `gorm:"uniqueIndex;not null"`
	Name          string  `gorm:"size:100"`
	Tier          string  `gorm:"size:20"` // Ultra | Pro | Standard
	CPU           string  `gorm:"size:150"`
	GPU           string  `gorm:"size:150"`
	RAM           string  `gorm:"size:50"`
	Storage       string  `gorm:"size:100"`
	Monitor       string  `gorm:"size:150"`
	MonitorSize   string  `gorm:"size:20"`
	RefreshRate   string  `gorm:"size:20"`
	Keyboard      string  `gorm:"size:100"`
	Mouse         string  `gorm:"size:100"`
	Headset       string  `gorm:"size:100"`
	HourlyRate    float64 `gorm:"not null"`
	Status        string  `gorm:"size:20;default:available"`
	ImageURL      string  `gorm:"size:255"`
	IsActive      bool    `gorm:"default:true"`
}
