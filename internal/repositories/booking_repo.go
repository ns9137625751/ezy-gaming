package repositories

import (
	"github.com/nishantshekhada/ezygaming/internal/models"
	"gorm.io/gorm"
)

type bookingRepo struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepo{db: db}
}

func (r *bookingRepo) Create(b *models.Booking) error {
	return r.db.Create(b).Error
}

func (r *bookingRepo) FindAll() ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Preload("TimeSlot").Order("created_at desc").Find(&bookings).Error
	return bookings, err
}

func (r *bookingRepo) FindByID(id uint) (*models.Booking, error) {
	var b models.Booking
	err := r.db.Preload("TimeSlot").First(&b, id).Error
	return &b, err
}

func (r *bookingRepo) UpdateStatus(id uint, status string) error {
	return r.db.Model(&models.Booking{}).Where("id = ?", id).Update("status", status).Error
}

func (r *bookingRepo) FindByDate(date string) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Preload("TimeSlot").
		Where("DATE(booking_date) = ?", date).
		Find(&bookings).Error
	return bookings, err
}

func (r *bookingRepo) CountAll() (int64, error) {
	var n int64
	err := r.db.Model(&models.Booking{}).Count(&n).Error
	return n, err
}

func (r *bookingRepo) CountToday() (int64, error) {
	var n int64
	err := r.db.Model(&models.Booking{}).Where("DATE(booking_date) = CURDATE()").Count(&n).Error
	return n, err
}

func (r *bookingRepo) CountByStatus(status string) (int64, error) {
	var n int64
	err := r.db.Model(&models.Booking{}).Where("status = ?", status).Count(&n).Error
	return n, err
}

func (r *bookingRepo) SumRevenue() (float64, error) {
	var total float64
	err := r.db.Model(&models.Booking{}).
		Where("status != ?", "cancelled").
		Select("COALESCE(SUM(total_amount), 0)").Scan(&total).Error
	return total, err
}

func (r *bookingRepo) FindRecent(limit int) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Preload("TimeSlot").Order("created_at desc").Limit(limit).Find(&bookings).Error
	return bookings, err
}

func (r *bookingRepo) FindByCustomerEmail(email string) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Preload("TimeSlot").
		Where("customer_email = ?", email).
		Order("created_at desc").
		Find(&bookings).Error
	return bookings, err
}
