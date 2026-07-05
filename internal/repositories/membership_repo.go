package repositories

import (
	"github.com/nishantshekhada/ezygaming/internal/models"
	"gorm.io/gorm"
)

type membershipPlanRepo struct {
	db *gorm.DB
}

func NewMembershipPlanRepository(db *gorm.DB) MembershipPlanRepository {
	return &membershipPlanRepo{db: db}
}

func (r *membershipPlanRepo) FindAll() ([]models.MembershipPlan, error) {
	var plans []models.MembershipPlan
	err := r.db.Order("sort_order asc").Find(&plans).Error
	return plans, err
}

func (r *membershipPlanRepo) FindActive() ([]models.MembershipPlan, error) {
	var plans []models.MembershipPlan
	err := r.db.Where("is_active = ?", true).Order("sort_order asc").Find(&plans).Error
	return plans, err
}

func (r *membershipPlanRepo) FindByID(id uint) (*models.MembershipPlan, error) {
	var p models.MembershipPlan
	err := r.db.First(&p, id).Error
	return &p, err
}

func (r *membershipPlanRepo) Create(p *models.MembershipPlan) error {
	return r.db.Create(p).Error
}

func (r *membershipPlanRepo) Update(p *models.MembershipPlan) error {
	return r.db.Save(p).Error
}
