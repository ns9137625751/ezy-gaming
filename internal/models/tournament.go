package models

import (
	"time"

	"gorm.io/gorm"
)

type Tournament struct {
	gorm.Model
	Title                string    `gorm:"size:200;not null"`
	GameName             string    `gorm:"size:100"`
	Platform             string    `gorm:"size:20"` // pc | ps5 | both
	Format               string    `gorm:"size:100"` // 1v1, 5v5, Battle Royale
	PrizePool            float64
	EntryFee             float64   `gorm:"default:0"`
	MaxTeams             int
	RegisteredTeams      int       `gorm:"default:0"`
	StartDate            time.Time
	RegistrationDeadline time.Time
	Status               string    `gorm:"size:20;default:upcoming"` // upcoming | ongoing | completed
	BannerURL            string    `gorm:"size:255"`
	Description          string    `gorm:"type:text"`
	Rules                string    `gorm:"type:text"`
	IsActive             bool      `gorm:"default:true"`
}

type TournamentRegistration struct {
	gorm.Model
	TournamentID  uint   `gorm:"not null"`
	Tournament    Tournament `gorm:"foreignKey:TournamentID"`
	TeamName      string `gorm:"size:100"`
	CaptainName   string `gorm:"size:100;not null"`
	CaptainEmail  string `gorm:"size:150;not null"`
	CaptainPhone  string `gorm:"size:20;not null"`
	MemberCount   int    `gorm:"default:1"`
	PaymentStatus string `gorm:"size:20;default:unpaid"`
}
