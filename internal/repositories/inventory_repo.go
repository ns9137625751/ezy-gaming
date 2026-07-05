package repositories

import (
	"github.com/nishantshekhada/ezygaming/internal/models"
	"gorm.io/gorm"
)

type inventoryRepo struct{ db *gorm.DB }

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepo{db}
}

func (r *inventoryRepo) FindAll() ([]models.Inventory, error) {
	var items []models.Inventory
	err := r.db.Order("created_at desc").Find(&items).Error
	return items, err
}

func (r *inventoryRepo) FindByID(id uint) (*models.Inventory, error) {
	var item models.Inventory
	err := r.db.First(&item, id).Error
	return &item, err
}

func (r *inventoryRepo) Create(item *models.Inventory) error {
	return r.db.Create(item).Error
}

func (r *inventoryRepo) Update(item *models.Inventory) error {
	return r.db.Save(item).Error
}

func (r *inventoryRepo) Delete(id uint) error {
	return r.db.Delete(&models.Inventory{}, id).Error
}

func (r *inventoryRepo) Count() (int64, error) {
	var n int64
	err := r.db.Model(&models.Inventory{}).Where("is_active = ?", true).Count(&n).Error
	return n, err
}
