package seo

import "html/template"

// Business constants — single source of truth for all structured data + meta tags.
const (
	SiteURL      = "https://www.ezygaming.in"
	SiteName     = "Ezy Gaming"
	BusinessName = "Ezy Gaming"
	Tagline      = "Ahmedabad's Premier Gaming Café"
	Street       = "Vishala Supreme, 5-6, Sardar Patel Ring Rd, near Torrent Power subStation"
	Locality     = "Nikol"
	City         = "Ahmedabad"
	Region       = "Gujarat"
	PostalCode   = "380049"
	Country      = "IN"
	Phone        = "+91-79-40123456"
	PhoneDisplay = "+91 79 4012 3456"
	Email        = "hello@ezygaming.in"
	Lat          = "23.0523"
	Lng          = "72.6416"
	MapsURL      = "https://maps.google.com/maps?vet=10CAAQoqAOahcKEwjoq_zt37OVAxUAAAAAHQAAAAAQDQ..i&pvq=Cg0vZy8xMXhjeGhnOWtqIhAKCmV6eSBnYW1pbmcQAhgD&lqi=CgplenkgZ2FtaW5nSLDM3MevvICACFocEAAQARgAGAEiCmV6eSBnYW1pbmcqBggCEAAQAZIBEHZpZGVvX2dhbWVfc3RvcmU&fvr=1&cs=1&um=1&ie=UTF-8&fb=1&gl=in&sa=X&ftid=0x395e878d47db93b3:0xf679bbe326ebc0f2"
	DefaultOGImg = "/static/img/og-default.jpg"
	TwitterHandle = "@ezygamingzone"
)

// Meta holds all per-page SEO fields.
type Meta struct {
	Title       string       // plain title, e.g. "Gaming PCs"
	Description string       // ≤160 chars
	Keywords    string       // comma-separated keywords
	Canonical   string       // full canonical URL, e.g. "https://…/gaming-pcs"
	OGType      string       // "website" | "article"
	OGImage     string       // absolute URL
	Breadcrumbs []Breadcrumb // nil = don't render breadcrumbs
	NoIndex     bool         // true only for thank-you / error pages
}

// Breadcrumb is one step in the breadcrumb trail.
type Breadcrumb struct {
	Name string
	URL  string
}

// Defaults fills any zero-value fields with site-wide defaults.
func (m *Meta) Defaults() *Meta {
	if m.OGType == "" {
		m.OGType = "website"
	}
	if m.OGImage == "" {
		m.OGImage = SiteURL + DefaultOGImg
	}
	if m.Canonical == "" {
		m.Canonical = SiteURL
	}
	return m
}

// FullTitle returns "Page | Ezy Gaming" or the site tagline if Title is empty.
func (m *Meta) FullTitle() string {
	if m.Title == "" {
		return BusinessName + " — " + Tagline
	}
	return m.Title + " | " + SiteName
}

// FullDescription returns the page description or a keyword-rich fallback.
func (m *Meta) FullDescription() string {
	if m.Description != "" {
		return m.Description
	}
	return "Ezy Gaming Zone — gaming café in Nikol, Ahmedabad. 20 high-end Gaming PCs, 20 PS5 consoles, esports tournaments, gaming memberships. Visit us in Nikol, Ahmedabad, Gujarat."
}

// PageData bundles Meta with optional pre-rendered safe JSON-LD blobs.
type PageData struct {
	Meta    *Meta
	JSONLDs []template.HTML // each entry is a full <script type="application/ld+json">…</script>
}
