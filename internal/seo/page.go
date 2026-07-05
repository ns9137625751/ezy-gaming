package seo

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

// PageConfig is the input used to build a gin.H for a page render.
type PageConfig struct {
	Title           string
	Description     string
	Keywords        string
	Path            string // e.g. "/gaming-pcs"
	OGType          string
	OGImage         string
	Breadcrumbs     []Breadcrumb
	NoIndex         bool
	ExtraJSONLDs    []template.HTML // page-specific JSON-LD (FAQ, Event, Service…)
	ExtraData       gin.H           // any additional controller data
}

// Build assembles the full gin.H that the renderer template expects.
// It merges SEO fields, global JSON-LD, and any controller-supplied extra data.
func Build(cfg PageConfig) gin.H {
	m := &Meta{
		Title:       cfg.Title,
		Description: cfg.Description,
		Keywords:    cfg.Keywords,
		Canonical:   SiteURL + cfg.Path,
		OGType:      cfg.OGType,
		OGImage:     cfg.OGImage,
		Breadcrumbs: cfg.Breadcrumbs,
		NoIndex:     cfg.NoIndex,
	}
	m.Defaults()

	data := gin.H{
		// Legacy key still used by some partial templates.
		"Title": cfg.Title,

		// SEO struct (accessed as .SEO.FullTitle, .SEO.Breadcrumbs, etc.)
		"SEO": m,

		// Global JSON-LD on every page.
		"LocalBusinessLD": LocalBusinessSchema(),
		"WebSiteLD":       WebSiteSchema(),

		// Breadcrumb JSON-LD (empty string if no breadcrumbs).
		"BreadcrumbLD": BreadcrumbSchema(cfg.Breadcrumbs),

		// Page-specific extra JSON-LD blobs (FAQPage, Event, Service…).
		"ExtraLDs": cfg.ExtraJSONLDs,
	}

	// Merge controller data.
	for k, v := range cfg.ExtraData {
		data[k] = v
	}

	return data
}

// HomeBreadcrumb is the standard home item — callers can prepend it.
var HomeBreadcrumb = Breadcrumb{Name: "Home", URL: "/"}
