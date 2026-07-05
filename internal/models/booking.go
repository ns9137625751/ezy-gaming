package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	BookingRef    string    `gorm:"uniqueIndex;size:20;not null"`
	CustomerName  string    `gorm:"size:100;not null"`
	CustomerEmail string    `gorm:"size:150;not null"`
	CustomerPhone string    `gorm:"size:20;not null"`
	StationType   string    `gorm:"size:20;not null"` // pc | ps4 | car_controller
	StationID     uint      `gorm:"not null"`
	BookingDate   time.Time `gorm:"not null"`
	TimeSlotID    *uint
	TimeSlot      TimeSlot `gorm:"foreignKey:TimeSlotID"`
	DurationHours int      `gorm:"default:1"`
	TotalAmount   float64  `gorm:"not null"`
	Status        string   `gorm:"size:20;default:pending"` // pending | confirmed | cancelled | completed
	PaymentStatus string   `gorm:"size:20;default:unpaid"`  // unpaid | paid
	Notes         string   `gorm:"type:text"`
}
