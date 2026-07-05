package controllers

import (
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/seo"
)

type SEOController struct{}

func NewSEOController() *SEOController { return &SEOController{} }

// Sitemap serves /sitemap.xml
func (h *SEOController) Sitemap(c *gin.Context) {
	sitemap := seo.GenerateSitemap()
	output, err := xml.MarshalIndent(sitemap, "", "  ")
	if err != nil {
		c.String(http.StatusInternalServerError, "sitemap error")
		return
	}
	c.Data(http.StatusOK, "application/xml; charset=utf-8", append([]byte(xml.Header), output...))
}

// RobotsTxt serves /robots.txt
func (h *SEOController) RobotsTxt(c *gin.Context) {
	content := `User-agent: *
Allow: /
Disallow: /static/
Disallow: /admin/

# Priority pages for crawlers
Sitemap: https://www.ezygaming.in/sitemap.xml

# Crawl delay for polite bots
Crawl-delay: 1
`
	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(content))
}

// HumansTxt serves /humans.txt (optional but good for AI crawlers)
func (h *SEOController) HumansTxt(c *gin.Context) {
	content := `/* TEAM */
Developer: Ezy Gaming Tech Team
Contact: hello@ezygaming.in
Location: Nikol, Ahmedabad, Gujarat, India

/* SITE */
Last update: 2025
Language: English, Hindi, Gujarati
Doctype: HTML5
IDE: Go, Tailwind CSS

/* ABOUT */
Ezy Gaming is Ahmedabad's premier gaming café located in Nikol.
We offer high-end Gaming PCs (RTX 4090), PS5 consoles on 4K displays,
esports tournaments, and gaming membership plans.
`
	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(content))
}

// SecurityTxt serves /.well-known/security.txt
func (h *SEOController) SecurityTxt(c *gin.Context) {
	content := `Contact: mailto:security@ezygaming.in
Expires: 2026-12-31T23:59:59Z
Preferred-Languages: en
`
	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(content))
}
