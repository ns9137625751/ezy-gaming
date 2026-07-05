package repositories

import (
	"github.com/nishantshekhada/ezygaming/internal/models"
	"gorm.io/gorm"
)

type ps5Repo struct{ db *gorm.DB }

func NewPS5Repository(db *gorm.DB) PS5Repository {
	return &ps5Repo{db: db}
}

func (r *ps5Repo) FindAll() ([]models.PS5Console, error) {
	var cs []models.PS5Console
	return cs, r.db.Order("station_number asc").Find(&cs).Error
}

func (r *ps5Repo) FindByID(id uint) (*models.PS5Console, error) {
	var c models.PS5Console
	return &c, r.db.First(&c, id).Error
}

func (r *ps5Repo) FindActive() ([]models.PS5Console, error) {
	var cs []models.PS5Console
	return cs, r.db.Where("is_active = ?", true).Order("station_number asc").Find(&cs).Error
}

func (r *ps5Repo) Create(c *models.PS5Console) error  { return r.db.Create(c).Error }
func (r *ps5Repo) Update(c *models.PS5Console) error  { return r.db.Save(c).Error }
func (r *ps5Repo) Delete(id uint) error               { return r.db.Delete(&models.PS5Console{}, id).Error }
