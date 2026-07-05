package repositories

import (
	"github.com/nishantshekhada/ezygaming/internal/models"
	"gorm.io/gorm"
)

type tournamentRepo struct {
	db *gorm.DB
}

func NewTournamentRepository(db *gorm.DB) TournamentRepository {
	return &tournamentRepo{db: db}
}

func (r *tournamentRepo) FindAll() ([]models.Tournament, error) {
	var ts []models.Tournament
	err := r.db.Order("start_date desc").Find(&ts).Error
	return ts, err
}

func (r *tournamentRepo) FindByID(id uint) (*models.Tournament, error) {
	var t models.Tournament
	err := r.db.First(&t, id).Error
	return &t, err
}

func (r *tournamentRepo) FindUpcoming(limit int) ([]models.Tournament, error) {
	var ts []models.Tournament
	err := r.db.Where("status = ? AND is_active = ?", "upcoming", true).
		Order("start_date asc").
		Limit(limit).
		Find(&ts).Error
	return ts, err
}

func (r *tournamentRepo) Create(t *models.Tournament) error {
	return r.db.Create(t).Error
}

func (r *tournamentRepo) Update(t *models.Tournament) error {
	return r.db.Save(t).Error
}

func (r *tournamentRepo) Delete(id uint) error {
	return r.db.Delete(&models.Tournament{}, id).Error
}
