package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type StringSlice []string

func (s StringSlice) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return string(b), err
}

func (s *StringSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan StringSlice")
	}
	return json.Unmarshal(bytes, s)
}

type MembershipPlan struct {
	gorm.Model
	Name         string      `gorm:"size:100;not null"`
	Slug         string      `gorm:"size:50;uniqueIndex"` // bronze, silver, gold, platinum
	DurationDays int         `gorm:"not null"`
	Price        float64     `gorm:"not null"`
	PCHours      int         `gorm:"default:0"`
	PS5Hours     int         `gorm:"default:0"`
	DiscountPct  int         `gorm:"default:0"`
	Benefits     StringSlice `gorm:"type:text"`
	Color        string      `gorm:"size:20"` // for UI theming: cyan, purple, orange, gold
	BadgeText    string      `gorm:"size:30"` // "Most Popular", "Best Value", etc.
	IsActive     bool        `gorm:"default:true"`
	SortOrder    int         `gorm:"default:0"`
}
