package repositories

import (
	"github.com/nishantshekhada/ezygaming/internal/models"
	"gorm.io/gorm"
)

type contactRepo struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepo{db: db}
}

func (r *contactRepo) Create(msg *models.ContactMessage) error {
	return r.db.Create(msg).Error
}

func (r *contactRepo) FindAll() ([]models.ContactMessage, error) {
	var msgs []models.ContactMessage
	err := r.db.Order("created_at desc").Find(&msgs).Error
	return msgs, err
}

func (r *contactRepo) MarkRead(id uint) error {
	return r.db.Model(&models.ContactMessage{}).Where("id = ?", id).Update("is_read", true).Error
}

func (r *contactRepo) CountUnread() (int64, error) {
	var n int64
	err := r.db.Model(&models.ContactMessage{}).Where("is_read = false").Count(&n).Error
	return n, err
}
