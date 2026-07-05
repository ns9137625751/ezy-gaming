package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/repositories"
	"github.com/nishantshekhada/ezygaming/internal/seo"
	"github.com/nishantshekhada/ezygaming/pkg/renderer"
)

type GalleryController struct {
	repo repositories.GalleryRepository
}

func NewGalleryController(repo repositories.GalleryRepository) *GalleryController {
	return &GalleryController{repo: repo}
}

func (h *GalleryController) Index(c *gin.Context) {
	images, _ := h.repo.FindAll()
	data := seo.Build(seo.PageConfig{
		Title:       "Gallery — Ezy Gaming Zone Ahmedabad",
		Description: "Browse photos of Ezy Gaming Zone — our setups, events, tournaments, and gaming café in Nikol, Ahmedabad.",
		Keywords:    "gaming cafe photos ahmedabad, ezy gaming gallery, gaming zone images nikol",
		Path:        "/gallery",
		Breadcrumbs: []seo.Breadcrumb{{Name: "Gallery"}},
		ExtraData:   gin.H{"Images": images},
	})
	renderer.Render(c, http.StatusOK, "gallery", data)
}
