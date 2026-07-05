package repositories

import (
	"github.com/nishantshekhada/ezygaming/internal/models"
	"gorm.io/gorm"
)

type timeSlotRepo struct{ db *gorm.DB }

func NewTimeSlotRepository(db *gorm.DB) TimeSlotRepository {
	return &timeSlotRepo{db: db}
}

func (r *timeSlotRepo) FindActive() ([]models.TimeSlot, error) {
	var slots []models.TimeSlot
	return slots, r.db.Where("is_active = ?", true).Order("start_time asc").Find(&slots).Error
}
