package database

import (
	"log"

	"github.com/nishantshekhada/ezygaming/internal/config"
	"github.com/nishantshekhada/ezygaming/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(cfg *config.Config) *gorm.DB {
	// Warn: shows only slow queries (>200ms) and errors — not every SELECT.
	// Change to logger.Info to see all SQL during debugging.
	db, err := gorm.Open(mysql.Open(cfg.DB.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")
	DB = db
	return db
}

func Migrate(db *gorm.DB) {
	// Run migrations silently — schema checks generate hundreds of lines.
	silentDB := db.Session(&gorm.Session{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	err := silentDB.AutoMigrate(
		&models.AdminUser{},
		&models.Customer{},
		&models.GamingPC{},
		&models.PS5Console{},
		&models.Game{},
		&models.TimeSlot{},
		&models.Booking{},
		&models.MembershipPlan{},
		&models.Tournament{},
		&models.TournamentRegistration{},
		&models.GalleryImage{},
		&models.SnackItem{},
		&models.FAQ{},
		&models.ContactMessage{},
		&models.Inventory{},
	)
	if err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	}
	log.Println("Database migrations completed")
}
