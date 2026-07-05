package admin

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/middlewares"
	"github.com/nishantshekhada/ezygaming/internal/models"
	"github.com/nishantshekhada/ezygaming/internal/repositories"
)

type InventoryController struct {
	repo        repositories.InventoryRepository
	contactRepo repositories.ContactRepository
}

func NewInventoryController(repo repositories.InventoryRepository, contactRepo repositories.ContactRepository) *InventoryController {
	return &InventoryController{repo, contactRepo}
}

func (h *InventoryController) Index(c *gin.Context) {
	user := middlewares.GetAdminUser(c)
	items, _ := h.repo.FindAll()
	unread, _ := h.contactRepo.CountUnread()
	renderAdmin(c, http.StatusOK, "inventory", gin.H{
		"User":           user,
		"Page":           "inventory",
		"Items":          items,
		"Flash":          c.Query("flash"),
		"UnreadContacts": unread,
	})
}

func (h *InventoryController) Create(c *gin.Context) {
	systemType := strings.TrimSpace(c.PostForm("system_type"))
	monitorSize := strings.TrimSpace(c.PostForm("monitor_size"))
	maxPeople, _ := strconv.Atoi(c.PostForm("max_people"))
	pricePerHr, _ := strconv.ParseFloat(c.PostForm("price_per_hr"), 64)

	if systemType == "" || pricePerHr <= 0 || maxPeople <= 0 {
		c.Redirect(http.StatusSeeOther, "/admin/inventory?flash=invalid")
		return
	}

	item := &models.Inventory{
		SystemType:  systemType,
		MonitorSize: monitorSize,
		MaxPeople:   maxPeople,
		PricePerHr:  pricePerHr,
		IsActive:    true,
	}
	if err := h.repo.Create(item); err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/inventory?flash=error")
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin/inventory?flash=created")
}

func (h *InventoryController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/inventory")
		return
	}

	item, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/inventory?flash=notfound")
		return
	}

	item.SystemType  = strings.TrimSpace(c.PostForm("system_type"))
	item.MonitorSize = strings.TrimSpace(c.PostForm("monitor_size"))
	item.MaxPeople, _ = strconv.Atoi(c.PostForm("max_people"))
	item.PricePerHr, _ = strconv.ParseFloat(c.PostForm("price_per_hr"), 64)
	item.IsActive = c.PostForm("is_active") == "1"

	if item.SystemType == "" || item.PricePerHr <= 0 || item.MaxPeople <= 0 {
		c.Redirect(http.StatusSeeOther, "/admin/inventory?flash=invalid")
		return
	}

	if err := h.repo.Update(item); err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/inventory?flash=error")
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin/inventory?flash=updated")
}

func (h *InventoryController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/inventory")
		return
	}
	h.repo.Delete(uint(id))
	c.Redirect(http.StatusSeeOther, "/admin/inventory?flash=deleted")
}
