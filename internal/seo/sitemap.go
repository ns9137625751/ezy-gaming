package seo

import (
	"encoding/xml"
	"time"
)

// URLSet is the root element of a sitemap.xml.
type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	URLs    []SitemapURL `xml:"url"`
}

// SitemapURL represents one <url> entry in the sitemap.
type SitemapURL struct {
	Loc        string  `xml:"loc"`
	LastMod    string  `xml:"lastmod,omitempty"`
	ChangeFreq string  `xml:"changefreq,omitempty"`
	Priority   float64 `xml:"priority"`
}

// GenerateSitemap builds a URLSet for all public pages.
func GenerateSitemap() *URLSet {
	today := time.Now().Format("2006-01-02")

	pages := []struct {
		path   string
		freq   string
		pri    float64
		static bool // false = updated daily
	}{
		{"/", "daily", 1.0, false},
		{"/gaming-pcs", "weekly", 0.9, false},
		{"/ps5-zone", "weekly", 0.9, false},
		{"/pricing", "weekly", 0.85, false},
		{"/memberships", "weekly", 0.85, false},
		{"/tournaments", "daily", 0.9, false},
		{"/contact", "monthly", 0.7, true},
		{"/book", "weekly", 0.95, false},
	}

	urls := make([]SitemapURL, 0, len(pages))
	for _, p := range pages {
		urls = append(urls, SitemapURL{
			Loc:        SiteURL + p.path,
			LastMod:    today,
			ChangeFreq: p.freq,
			Priority:   p.pri,
		})
	}

	return &URLSet{
		XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  urls,
	}
}
