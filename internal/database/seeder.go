package database

import (
	"log"

	"github.com/nishantshekhada/ezygaming/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	seedAdminUsers(db)
	seedGamingPCs(db)
	seedPS5Consoles(db)
	seedMembershipPlans(db)
	seedTimeSlots(db)
	log.Println("Database seeded successfully")
}

func seedAdminUsers(db *gorm.DB) {
	var count int64
	db.Model(&models.AdminUser{}).Count(&count)
	if count > 0 {
		return
	}

	users := []struct {
		name, email, password, role string
	}{
		{"Admin", "admin@ezygaming.in", "admin@123", "admin"},
		{"Staff", "staff@ezygaming.in", "staff@123", "gamer"},
	}

	for _, u := range users {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("bcrypt error for %s: %v", u.email, err)
			continue
		}
		db.Create(&models.AdminUser{
			Name:     u.name,
			Email:    u.email,
			Password: string(hash),
			Role:     u.role,
			IsActive: true,
		})
	}
	log.Println("Admin users seeded — admin@ezygaming.in / admin@123 | staff@ezygaming.in / staff@123")
}

func seedGamingPCs(db *gorm.DB) {
	var count int64
	db.Model(&models.GamingPC{}).Count(&count)
	if count > 0 {
		return
	}

	type tier struct {
		gpu, cpu, ram, storage, monitor, monitorSize, refreshRate, keyboard, mouse, headset, name, tierName string
		rate                                                                                                  float64
	}
	tiers := []tier{
		{"RTX 4090 24GB", "Intel i9-13900K", "32GB DDR5", "2TB NVMe SSD", "27\" IPS 240Hz", "27\"", "240Hz", "HyperX Alloy Origins", "Logitech G Pro X", "HyperX Cloud Alpha", "Ultra", "Ultra", 80},
		{"RTX 4080 16GB", "Intel i7-13700K", "32GB DDR5", "1TB NVMe SSD", "27\" IPS 165Hz", "27\"", "165Hz", "Corsair K70 RGB", "Razer DeathAdder V3", "SteelSeries Arctis 7", "Pro", "Pro", 70},
		{"RTX 3090 24GB", "Intel i7-12700K", "16GB DDR4", "1TB SSD", "24\" IPS 144Hz", "24\"", "144Hz", "Redragon K552", "Logitech G402", "Corsair HS50", "Standard", "Standard", 60},
	}

	pcs := []models.GamingPC{}
	for i := 1; i <= 20; i++ {
		var t tier
		switch {
		case i <= 8:
			t = tiers[0]
		case i <= 14:
			t = tiers[1]
		default:
			t = tiers[2]
		}
		pcs = append(pcs, models.GamingPC{
			StationNumber: i,
			Name:          t.name + " Station",
			Tier:          t.tierName,
			CPU:           t.cpu,
			GPU:           t.gpu,
			RAM:           t.ram,
			Storage:       t.storage,
			Monitor:       t.monitor,
			MonitorSize:   t.monitorSize,
			RefreshRate:   t.refreshRate,
			Keyboard:      t.keyboard,
			Mouse:         t.mouse,
			Headset:       t.headset,
			HourlyRate:    t.rate,
			Status:        "available",
			IsActive:      true,
		})
	}
	db.Create(&pcs)
}

func seedPS5Consoles(db *gorm.DB) {
	var count int64
	db.Model(&models.PS5Console{}).Count(&count)
	if count > 0 {
		return
	}

	consoles := []models.PS5Console{}
	for i := 1; i <= 20; i++ {
		tvSize := "55\" 4K OLED"
		rate := 60.0
		name := "Premium Station"
		tierName := "Premium"
		if i > 10 {
			tvSize = "50\" 4K QLED"
			rate = 50.0
			name = "Standard Station"
			tierName = "Standard"
		}
		consoles = append(consoles, models.PS5Console{
			StationNumber:   i,
			Name:            name,
			Tier:            tierName,
			TVSize:          tvSize,
			ControllerCount: 2,
			HourlyRate:      rate,
			Status:          "available",
			IsActive:        true,
		})
	}
	db.Create(&consoles)
}

func seedMembershipPlans(db *gorm.DB) {
	var count int64
	db.Model(&models.MembershipPlan{}).Count(&count)
	if count > 0 {
		return
	}

	plans := []models.MembershipPlan{
		{
			Name: "Bronze", Slug: "bronze", DurationDays: 30, Price: 499,
			PCHours: 5, PS5Hours: 0, DiscountPct: 5, SortOrder: 1,
			Color: "orange", BadgeText: "",
			Benefits: models.StringSlice{"5 PC hours included", "5% off on sessions", "Member badge", "Access to all games"},
			IsActive: true,
		},
		{
			Name: "Silver", Slug: "silver", DurationDays: 30, Price: 799,
			PCHours: 10, PS5Hours: 0, DiscountPct: 10, SortOrder: 2,
			Color: "gray", BadgeText: "",
			Benefits: models.StringSlice{"10 PC hours included", "10% off on sessions", "Tournament priority access", "Member badge + profile"},
			IsActive: true,
		},
		{
			Name: "Gold", Slug: "gold", DurationDays: 30, Price: 1499,
			PCHours: 25, PS5Hours: 10, DiscountPct: 20, SortOrder: 3,
			Color: "cyan", BadgeText: "Most Popular",
			Benefits: models.StringSlice{"25 PC + 10 PS5 hours", "20% off on all sessions", "1 free tournament entry/month", "Snack voucher ₹200/month", "Priority seat booking"},
			IsActive: true,
		},
		{
			Name: "Platinum", Slug: "platinum", DurationDays: 30, Price: 2499,
			PCHours: 0, PS5Hours: 0, DiscountPct: 30, SortOrder: 4,
			Color: "purple", BadgeText: "Best Value",
			Benefits: models.StringSlice{"Unlimited PC + PS5 hours", "30% off on everything", "Free unlimited tournaments", "Reserved VIP station", "Snack voucher ₹500/month", "Dedicated account manager"},
			IsActive: true,
		},
	}
	db.Create(&plans)
}

func seedTimeSlots(db *gorm.DB) {
	var count int64
	db.Model(&models.TimeSlot{}).Count(&count)
	if count > 0 {
		return
	}

	slots := []struct{ start, end, label string }{
		{"00:00", "01:00", "12:00 AM – 01:00 AM"},
		{"01:00", "02:00", "01:00 AM – 02:00 AM"},
		{"02:00", "03:00", "02:00 AM – 03:00 AM"},
		{"03:00", "04:00", "03:00 AM – 04:00 AM"},
		{"04:00", "05:00", "04:00 AM – 05:00 AM"},
		{"05:00", "06:00", "05:00 AM – 06:00 AM"},
		{"06:00", "07:00", "06:00 AM – 07:00 AM"},
		{"07:00", "08:00", "07:00 AM – 08:00 AM"},
		{"08:00", "09:00", "08:00 AM – 09:00 AM"},
		{"09:00", "10:00", "09:00 AM – 10:00 AM"},
		{"10:00", "11:00", "10:00 AM – 11:00 AM"},
		{"11:00", "12:00", "11:00 AM – 12:00 PM"},
		{"12:00", "13:00", "12:00 PM – 01:00 PM"},
		{"13:00", "14:00", "01:00 PM – 02:00 PM"},
		{"14:00", "15:00", "02:00 PM – 03:00 PM"},
		{"15:00", "16:00", "03:00 PM – 04:00 PM"},
		{"16:00", "17:00", "04:00 PM – 05:00 PM"},
		{"17:00", "18:00", "05:00 PM – 06:00 PM"},
		{"18:00", "19:00", "06:00 PM – 07:00 PM"},
		{"19:00", "20:00", "07:00 PM – 08:00 PM"},
		{"20:00", "21:00", "08:00 PM – 09:00 PM"},
		{"21:00", "22:00", "09:00 PM – 10:00 PM"},
		{"22:00", "23:00", "10:00 PM – 11:00 PM"},
		{"23:00", "00:00", "11:00 PM – 12:00 AM"},
	}

	ts := []models.TimeSlot{}
	for _, s := range slots {
		ts = append(ts, models.TimeSlot{Label: s.label, StartTime: s.start, EndTime: s.end, IsActive: true})
	}
	db.Create(&ts)
}
