package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/models"
	"github.com/nishantshekhada/ezygaming/internal/repositories"
	"github.com/nishantshekhada/ezygaming/internal/seo"
	"github.com/nishantshekhada/ezygaming/pkg/renderer"
)

type ContactController struct {
	repo repositories.ContactRepository
}

func NewContactController(repo repositories.ContactRepository) *ContactController {
	return &ContactController{repo: repo}
}

func (h *ContactController) Index(c *gin.Context) {
	data := seo.Build(seo.PageConfig{
		Title:       "Contact Ezy Gaming — Gaming Café in Nikol, Ahmedabad",
		Description: "Contact Ezy Gaming in Nikol, Ahmedabad. Call +91 79 4012 3456, email hello@ezygaming.in, or visit us at Vishala Supreme, Sardar Patel Ring Rd, Nikol. Open 24/7, every day of the year.",
		Keywords:    "contact gaming cafe ahmedabad, ezy gaming contact, gaming zone nikol address, gaming cafe ahmedabad phone number",
		Path:        "/contact",
		Breadcrumbs: []seo.Breadcrumb{{Name: "Contact"}},
		NoIndex:     false,
	})

	renderer.Render(c, http.StatusOK, "contact", data)
}

func (h *ContactController) Submit(c *gin.Context) {
	var req struct {
		Name    string `json:"name"    binding:"required"`
		Email   string `json:"email"   binding:"required,email"`
		Phone   string `json:"phone"`
		Subject string `json:"subject"`
		Message string `json:"message" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Please fill all required fields."})
		return
	}

	msg := &models.ContactMessage{
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Subject: req.Subject,
		Message: req.Message,
	}

	if err := h.repo.Create(msg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Could not save your message. Please try again."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Message sent! We'll get back to you within 24 hours."})
}
