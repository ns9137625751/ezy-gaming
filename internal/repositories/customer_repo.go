package repositories

import (
	"github.com/nishantshekhada/ezygaming/internal/models"
	"gorm.io/gorm"
)

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepo{db: db}
}

func (r *customerRepo) FindByEmail(email string) (*models.Customer, error) {
	var c models.Customer
	err := r.db.Where("email = ?", email).First(&c).Error
	return &c, err
}

func (r *customerRepo) FindByID(id uint) (*models.Customer, error) {
	var c models.Customer
	err := r.db.First(&c, id).Error
	return &c, err
}

func (r *customerRepo) Create(c *models.Customer) error {
	return r.db.Create(c).Error
}
