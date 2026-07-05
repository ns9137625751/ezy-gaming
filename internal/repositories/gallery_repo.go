package repositories

import (
	"github.com/nishantshekhada/ezygaming/internal/models"
	"gorm.io/gorm"
)

type galleryRepo struct {
	db *gorm.DB
}

func NewGalleryRepository(db *gorm.DB) GalleryRepository {
	return &galleryRepo{db: db}
}

func (r *galleryRepo) FindAll() ([]models.GalleryImage, error) {
	var imgs []models.GalleryImage
	err := r.db.Order("sort_order asc").Find(&imgs).Error
	return imgs, err
}

func (r *galleryRepo) FindFeatured(limit int) ([]models.GalleryImage, error) {
	var imgs []models.GalleryImage
	err := r.db.Where("is_featured = ?", true).Order("sort_order asc").Limit(limit).Find(&imgs).Error
	return imgs, err
}

func (r *galleryRepo) FindByCategory(category string) ([]models.GalleryImage, error) {
	var imgs []models.GalleryImage
	err := r.db.Where("category = ?", category).Order("sort_order asc").Find(&imgs).Error
	return imgs, err
}

func (r *galleryRepo) Create(img *models.GalleryImage) error {
	return r.db.Create(img).Error
}

func (r *galleryRepo) Delete(id uint) error {
	return r.db.Delete(&models.GalleryImage{}, id).Error
}
