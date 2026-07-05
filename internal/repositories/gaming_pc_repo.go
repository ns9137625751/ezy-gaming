package repositories

import (
	"github.com/nishantshekhada/ezygaming/internal/models"
	"gorm.io/gorm"
)

type gamingPCRepo struct {
	db *gorm.DB
}

func NewGamingPCRepository(db *gorm.DB) GamingPCRepository {
	return &gamingPCRepo{db: db}
}

func (r *gamingPCRepo) FindAll() ([]models.GamingPC, error) {
	var pcs []models.GamingPC
	err := r.db.Order("station_number asc").Find(&pcs).Error
	return pcs, err
}

func (r *gamingPCRepo) FindByID(id uint) (*models.GamingPC, error) {
	var pc models.GamingPC
	err := r.db.First(&pc, id).Error
	return &pc, err
}

func (r *gamingPCRepo) FindActive() ([]models.GamingPC, error) {
	var pcs []models.GamingPC
	err := r.db.Where("is_active = ?", true).Order("station_number asc").Find(&pcs).Error
	return pcs, err
}

func (r *gamingPCRepo) Create(pc *models.GamingPC) error {
	return r.db.Create(pc).Error
}

func (r *gamingPCRepo) Update(pc *models.GamingPC) error {
	return r.db.Save(pc).Error
}

func (r *gamingPCRepo) Delete(id uint) error {
	return r.db.Delete(&models.GamingPC{}, id).Error
}
