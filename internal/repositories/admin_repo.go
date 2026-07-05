package repositories

import (
	"github.com/nishantshekhada/ezygaming/internal/models"
	"gorm.io/gorm"
)

type adminUserRepo struct {
	db *gorm.DB
}

func NewAdminUserRepository(db *gorm.DB) AdminUserRepository {
	return &adminUserRepo{db: db}
}

func (r *adminUserRepo) FindByEmail(email string) (*models.AdminUser, error) {
	var u models.AdminUser
	err := r.db.Where("email = ? AND is_active = true", email).First(&u).Error
	return &u, err
}

func (r *adminUserRepo) FindByID(id uint) (*models.AdminUser, error) {
	var u models.AdminUser
	err := r.db.First(&u, id).Error
	return &u, err
}
