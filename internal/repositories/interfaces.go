package repositories

import "github.com/nishantshekhada/ezygaming/internal/models"

type AdminUserRepository interface {
	FindByEmail(email string) (*models.AdminUser, error)
	FindByID(id uint) (*models.AdminUser, error)
}

type GamingPCRepository interface {
	FindAll() ([]models.GamingPC, error)
	FindByID(id uint) (*models.GamingPC, error)
	FindActive() ([]models.GamingPC, error)
	Create(pc *models.GamingPC) error
	Update(pc *models.GamingPC) error
	Delete(id uint) error
}

type PS5Repository interface {
	FindAll() ([]models.PS5Console, error)
	FindByID(id uint) (*models.PS5Console, error)
	FindActive() ([]models.PS5Console, error)
	Create(ps5 *models.PS5Console) error
	Update(ps5 *models.PS5Console) error
	Delete(id uint) error
}

type TournamentRepository interface {
	FindAll() ([]models.Tournament, error)
	FindByID(id uint) (*models.Tournament, error)
	FindUpcoming(limit int) ([]models.Tournament, error)
	Create(t *models.Tournament) error
	Update(t *models.Tournament) error
	Delete(id uint) error
}

type MembershipPlanRepository interface {
	FindAll() ([]models.MembershipPlan, error)
	FindActive() ([]models.MembershipPlan, error)
	FindByID(id uint) (*models.MembershipPlan, error)
	Create(p *models.MembershipPlan) error
	Update(p *models.MembershipPlan) error
}

type GalleryRepository interface {
	FindAll() ([]models.GalleryImage, error)
	FindFeatured(limit int) ([]models.GalleryImage, error)
	FindByCategory(category string) ([]models.GalleryImage, error)
	Create(img *models.GalleryImage) error
	Delete(id uint) error
}

type ContactRepository interface {
	Create(msg *models.ContactMessage) error
	FindAll() ([]models.ContactMessage, error)
	MarkRead(id uint) error
	CountUnread() (int64, error)
}

type TimeSlotRepository interface {
	FindActive() ([]models.TimeSlot, error)
}

type BookingRepository interface {
	Create(b *models.Booking) error
	FindAll() ([]models.Booking, error)
	FindByID(id uint) (*models.Booking, error)
	UpdateStatus(id uint, status string) error
	FindByDate(date string) ([]models.Booking, error)
	CountAll() (int64, error)
	CountToday() (int64, error)
	CountByStatus(status string) (int64, error)
	SumRevenue() (float64, error)
	FindRecent(limit int) ([]models.Booking, error)
	FindByCustomerEmail(email string) ([]models.Booking, error)
}

type CustomerRepository interface {
	FindByEmail(email string) (*models.Customer, error)
	FindByID(id uint) (*models.Customer, error)
	Create(c *models.Customer) error
}

type InventoryRepository interface {
	FindAll() ([]models.Inventory, error)
	FindByID(id uint) (*models.Inventory, error)
	Create(item *models.Inventory) error
	Update(item *models.Inventory) error
	Delete(id uint) error
	Count() (int64, error)
}

type FAQRepository interface {
	FindAll() ([]models.FAQ, error)
	FindActive() ([]models.FAQ, error)
	FindByCategory(category string) ([]models.FAQ, error)
}

