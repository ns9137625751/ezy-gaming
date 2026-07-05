package repositories

import (
	"github.com/nishantshekhada/ezygaming/internal/models"
	"gorm.io/gorm"
)

type faqRepo struct{ db *gorm.DB }

func NewFAQRepository(db *gorm.DB) FAQRepository {
	return &faqRepo{db: db}
}

func (r *faqRepo) FindAll() ([]models.FAQ, error) {
	var faqs []models.FAQ
	err := r.db.Order("sort_order asc, created_at asc").Find(&faqs).Error
	return faqs, err
}

func (r *faqRepo) FindActive() ([]models.FAQ, error) {
	var faqs []models.FAQ
	err := r.db.Where("is_active = ?", true).Order("sort_order asc, created_at asc").Find(&faqs).Error
	return faqs, err
}

func (r *faqRepo) FindByCategory(category string) ([]models.FAQ, error) {
	var faqs []models.FAQ
	err := r.db.Where("category = ? AND is_active = ?", category, true).Order("sort_order asc").Find(&faqs).Error
	return faqs, err
}
