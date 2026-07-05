package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	SystemType  string  `gorm:"size:30;not null"`  // "pc", "ps4", "car_controller"
	MonitorSize string  `gorm:"size:50"`
	MaxPeople   int     `gorm:"not null;default:1"`
	PricePerHr  float64 `gorm:"not null"`
	IsActive    bool    `gorm:"default:true"`
}
