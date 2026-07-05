package admin

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/middlewares"
	"github.com/nishantshekhada/ezygaming/internal/models"
	"github.com/nishantshekhada/ezygaming/internal/repositories"
)

type DashboardController struct {
	bookingRepo   repositories.BookingRepository
	contactRepo   repositories.ContactRepository
	inventoryRepo repositories.InventoryRepository
}

func NewDashboardController(
	bookingRepo repositories.BookingRepository,
	contactRepo repositories.ContactRepository,
	inventoryRepo repositories.InventoryRepository,
) *DashboardController {
	return &DashboardController{bookingRepo, contactRepo, inventoryRepo}
}

func (h *DashboardController) Dashboard(c *gin.Context) {
	user := middlewares.GetAdminUser(c)

	totalBookings, _ := h.bookingRepo.CountAll()
	todayBookings, _ := h.bookingRepo.CountToday()
	pendingBookings, _ := h.bookingRepo.CountByStatus("pending")
	recent, _ := h.bookingRepo.FindRecent(5)
	inventoryCount, _ := h.inventoryRepo.Count()
	allItems, _ := h.inventoryRepo.FindAll()

	var activeInventory []models.Inventory
	for _, item := range allItems {
		if item.IsActive {
			activeInventory = append(activeInventory, item)
		}
	}

	data := gin.H{
		"User":            user,
		"Page":            "dashboard",
		"TotalBookings":   totalBookings,
		"TodayBookings":   todayBookings,
		"PendingBookings": pendingBookings,
		"RecentBookings":  recent,
		"InventoryCount":  inventoryCount,
		"Inventory":       activeInventory,
		"Flash":           c.Query("flash"),
		"Today":           time.Now().Format("2006-01-02"),
	}

	if user.Role == "admin" {
		revenue, _ := h.bookingRepo.SumRevenue()
		unread, _ := h.contactRepo.CountUnread()
		data["Revenue"] = revenue
		data["UnreadContacts"] = unread
	}

	renderAdmin(c, http.StatusOK, "dashboard", data)
}

func (h *DashboardController) CreateBooking(c *gin.Context) {
	name := strings.TrimSpace(c.PostForm("name"))
	phone := strings.TrimSpace(c.PostForm("phone"))
	email := strings.TrimSpace(c.PostForm("email"))
	stationID, _ := strconv.ParseUint(c.PostForm("station_id"), 10, 64)
	duration, _ := strconv.Atoi(c.PostForm("duration"))
	amount, _ := strconv.ParseFloat(c.PostForm("amount"), 64)
	notes := strings.TrimSpace(c.PostForm("notes"))
	dateStr := c.PostForm("booking_date")

	if name == "" || phone == "" || stationID == 0 || duration < 1 || amount <= 0 || dateStr == "" {
		c.Redirect(http.StatusSeeOther, "/admin/dashboard?flash=invalid")
		return
	}

	bookDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/dashboard?flash=invalid")
		return
	}

	item, err := h.inventoryRepo.FindByID(uint(stationID))
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/dashboard?flash=invalid")
		return
	}

	ref := fmt.Sprintf("EZY-%d-%05d", time.Now().Year(), rand.Intn(99999)+1)

	booking := &models.Booking{
		BookingRef:    ref,
		CustomerName:  name,
		CustomerEmail: email,
		CustomerPhone: phone,
		StationType:   item.SystemType,
		StationID:     uint(stationID),
		BookingDate:   bookDate,
		DurationHours: duration,
		TotalAmount:   amount,
		Status:        "confirmed",
		PaymentStatus: "unpaid",
		Notes:         notes,
	}

	if err := h.bookingRepo.Create(booking); err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/dashboard?flash=error")
		return
	}

	c.Redirect(http.StatusSeeOther, "/admin/dashboard?flash=booked")
}

func (h *DashboardController) Bookings(c *gin.Context) {
	user := middlewares.GetAdminUser(c)
	filter := c.DefaultQuery("status", "all")

	bookings, _ := h.bookingRepo.FindAll()

	// Client-side filter applied in template; pass filter string for active tab
	var unread int64
	if user.Role == "admin" {
		unread, _ = h.contactRepo.CountUnread()
	}

	renderAdmin(c, http.StatusOK, "bookings", gin.H{
		"User":           user,
		"Page":           "bookings",
		"Bookings":       bookings,
		"Filter":         filter,
		"UnreadContacts": unread,
	})
}

func (h *DashboardController) UpdateBookingStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	status := c.PostForm("status")
	allowed := map[string]bool{"pending": true, "confirmed": true, "cancelled": true, "completed": true}
	if !allowed[status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status"})
		return
	}
	if err := h.bookingRepo.UpdateStatus(uint(id), status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin/bookings")
}

func (h *DashboardController) Contacts(c *gin.Context) {
	user := middlewares.GetAdminUser(c)
	msgs, _ := h.contactRepo.FindAll()
	unread, _ := h.contactRepo.CountUnread()

	renderAdmin(c, http.StatusOK, "contacts", gin.H{
		"User":           user,
		"Page":           "contacts",
		"Messages":       msgs,
		"UnreadContacts": unread,
	})
}

func (h *DashboardController) MarkContactRead(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/contacts")
		return
	}
	h.contactRepo.MarkRead(uint(id))
	c.Redirect(http.StatusSeeOther, "/admin/contacts")
}
