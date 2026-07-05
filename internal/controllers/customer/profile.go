package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/auth"
	"github.com/nishantshekhada/ezygaming/internal/repositories"
)

type ProfileController struct {
	custRepo    repositories.CustomerRepository
	bookingRepo repositories.BookingRepository
}

func NewProfileController(custRepo repositories.CustomerRepository, bookingRepo repositories.BookingRepository) *ProfileController {
	return &ProfileController{custRepo: custRepo, bookingRepo: bookingRepo}
}

func (h *ProfileController) Me(c *gin.Context) {
	session := c.MustGet("customer").(*auth.SessionUser)

	cust, err := h.custRepo.FindByID(session.ID)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	bookings, _ := h.bookingRepo.FindByCustomerEmail(cust.Email)

	// Compute stats in controller — keeps the template logic-free
	confirmed := 0
	totalSpent := 0.0
	for _, b := range bookings {
		if b.Status == "confirmed" || b.Status == "completed" {
			confirmed++
		}
		if b.Status != "cancelled" {
			totalSpent += b.TotalAmount
		}
	}

	renderPage(c, http.StatusOK, "profile", gin.H{
		"Customer":   cust,
		"Bookings":   bookings,
		"Confirmed":  confirmed,
		"TotalSpent": totalSpent,
		"Page":       "profile",
	})
}
