package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/models"
	"github.com/nishantshekhada/ezygaming/internal/repositories"
	"github.com/nishantshekhada/ezygaming/internal/seo"
	"github.com/nishantshekhada/ezygaming/pkg/renderer"
)

type BookingController struct {
	inventoryRepo repositories.InventoryRepository
	slotRepo      repositories.TimeSlotRepository
	bookingRepo   repositories.BookingRepository
}

func NewBookingController(
	inventoryRepo repositories.InventoryRepository,
	slotRepo repositories.TimeSlotRepository,
	bookingRepo repositories.BookingRepository,
) *BookingController {
	return &BookingController{inventoryRepo, slotRepo, bookingRepo}
}

func (h *BookingController) Index(c *gin.Context) {
	allItems, _ := h.inventoryRepo.FindAll()
	slots, _ := h.slotRepo.FindActive()

	var pcs, ps4s, cars []models.Inventory
	for _, item := range allItems {
		if !item.IsActive {
			continue
		}
		switch item.SystemType {
		case "pc":
			pcs = append(pcs, item)
		case "ps4":
			ps4s = append(ps4s, item)
		case "car_controller":
			cars = append(cars, item)
		}
	}

	ref := c.Query("ref")
	success := c.Query("success") == "true"

	data := seo.Build(seo.PageConfig{
		Title:       "Book a Gaming Session in Ahmedabad — Ezy Gaming Zone",
		Description: "Book a Gaming PC, PS4 or Car Simulator session online at Ezy Gaming Zone, Nikol Ahmedabad. Choose your station, pick a time slot, and pay at the venue.",
		Keywords:    "book gaming cafe ahmedabad, gaming session booking ahmedabad, book gaming pc nikol, reserve ps4 ahmedabad, online gaming booking gujarat",
		Path:        "/book",
		Breadcrumbs: []seo.Breadcrumb{{Name: "Book Session"}},
		ExtraData: gin.H{
			"PCs":     pcs,
			"PS4s":    ps4s,
			"Cars":    cars,
			"Slots":   slots,
			"Success": success,
			"BookRef": ref,
		},
	})

	renderer.Render(c, http.StatusOK, "booking", data)
}

func (h *BookingController) Submit(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	stationType := c.PostForm("station_type")
	stationID, _ := strconv.ParseUint(c.PostForm("station_id"), 10, 64)
	slotID, _ := strconv.ParseUint(c.PostForm("slot_id"), 10, 64)
	duration, _ := strconv.Atoi(c.PostForm("duration"))
	notes := c.PostForm("notes")
	dateStr := c.PostForm("booking_date")

	if name == "" || email == "" || phone == "" || stationType == "" || stationID == 0 || slotID == 0 || duration < 1 || dateStr == "" {
		c.Redirect(http.StatusSeeOther, "/book?error=missing_fields")
		return
	}

	bookDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/book?error=invalid_date")
		return
	}

	item, err := h.inventoryRepo.FindByID(uint(stationID))
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/book?error=station_not_found")
		return
	}

	ref := fmt.Sprintf("EZY-%d-%05d", time.Now().Year(), rand.Intn(99999)+1)

	booking := &models.Booking{
		BookingRef:    ref,
		CustomerName:  name,
		CustomerEmail: email,
		CustomerPhone: phone,
		StationType:   stationType,
		StationID:     uint(stationID),
		BookingDate:   bookDate,
		TimeSlotID:    func(v uint) *uint { return &v }(uint(slotID)),
		DurationHours: duration,
		TotalAmount:   item.PricePerHr * float64(duration),
		Status:        "pending",
		PaymentStatus: "unpaid",
		Notes:         notes,
	}

	if err := h.bookingRepo.Create(booking); err != nil {
		c.Redirect(http.StatusSeeOther, "/book?error=save_failed")
		return
	}

	c.Redirect(http.StatusSeeOther, "/book?success=true&ref="+ref)
}
