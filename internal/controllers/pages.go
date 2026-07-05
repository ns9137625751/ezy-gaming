package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/seo"
	"github.com/nishantshekhada/ezygaming/pkg/renderer"
)

type PagesController struct{}

func NewPagesController() *PagesController { return &PagesController{} }

func (h *PagesController) About(c *gin.Context) {
	data := seo.Build(seo.PageConfig{
		Title:       "About Us — Ezy Gaming Zone Ahmedabad",
		Description: "Learn about Ezy Gaming Zone — Ahmedabad's premier gaming café in Nikol, our story, mission, and the team behind the experience.",
		Keywords:    "about ezy gaming ahmedabad, gaming cafe story nikol, gaming zone team ahmedabad",
		Path:        "/about",
		Breadcrumbs: []seo.Breadcrumb{{Name: "About Us"}},
		ExtraData:   gin.H{},
	})
	renderer.Render(c, http.StatusOK, "about", data)
}

func (h *PagesController) Privacy(c *gin.Context) {
	data := seo.Build(seo.PageConfig{
		Title:       "Privacy Policy — Ezy Gaming Zone",
		Description: "Privacy policy for Ezy Gaming Zone. How we collect, use, and protect your personal information.",
		Keywords:    "ezy gaming privacy policy",
		Path:        "/privacy",
		Breadcrumbs: []seo.Breadcrumb{{Name: "Privacy Policy"}},
		NoIndex:     true,
		ExtraData:   gin.H{},
	})
	renderer.Render(c, http.StatusOK, "privacy", data)
}

func (h *PagesController) Terms(c *gin.Context) {
	data := seo.Build(seo.PageConfig{
		Title:       "Terms of Service — Ezy Gaming Zone",
		Description: "Terms of service for Ezy Gaming Zone. Rules and conditions for using our gaming café and services.",
		Keywords:    "ezy gaming terms of service",
		Path:        "/terms",
		Breadcrumbs: []seo.Breadcrumb{{Name: "Terms of Service"}},
		NoIndex:     true,
		ExtraData:   gin.H{},
	})
	renderer.Render(c, http.StatusOK, "terms", data)
}
