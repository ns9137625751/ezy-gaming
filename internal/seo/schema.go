package seo

import (
	"encoding/json"
	"fmt"
	"html/template"
	"time"
)

// jsonLD wraps a marshalled value in a safe <script> block.
func jsonLD(v any) template.HTML {
	b, _ := json.MarshalIndent(v, "", "  ")
	return template.HTML(fmt.Sprintf(`<script type="application/ld+json">%s</script>`, b))
}

// LocalBusinessSchema returns the JSON-LD for the business (injected on every page).
func LocalBusinessSchema() template.HTML {
	type geo struct {
		Type      string `json:"@type"`
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	}
	type address struct {
		Type            string `json:"@type"`
		StreetAddress   string `json:"streetAddress"`
		AddressLocality string `json:"addressLocality"`
		AddressRegion   string `json:"addressRegion"`
		PostalCode      string `json:"postalCode"`
		AddressCountry  string `json:"addressCountry"`
	}
	type openingHours struct {
		Type        string   `json:"@type"`
		DayOfWeek   []string `json:"dayOfWeek"`
		Opens       string   `json:"opens"`
		Closes      string   `json:"closes"`
	}
	type aggregateRating struct {
		Type        string  `json:"@type"`
		RatingValue float64 `json:"ratingValue"`
		ReviewCount int     `json:"reviewCount"`
		BestRating  int     `json:"bestRating"`
	}
	type lb struct {
		Context         string            `json:"@context"`
		Type            []string          `json:"@type"`
		Name            string            `json:"name"`
		AlternateName   string            `json:"alternateName"`
		Description     string            `json:"description"`
		URL             string            `json:"url"`
		Telephone       string            `json:"telephone"`
		Email           string            `json:"email"`
		Image           string            `json:"image"`
		Logo            string            `json:"logo"`
		Address         address           `json:"address"`
		Geo             geo               `json:"geo"`
		OpeningHoursSpecification []openingHours `json:"openingHoursSpecification"`
		PriceRange      string            `json:"priceRange"`
		ServesCuisine   []string          `json:"amenityFeature,omitempty"`
		AggregateRating aggregateRating   `json:"aggregateRating"`
		SameAs          []string          `json:"sameAs"`
		Keywords        string            `json:"keywords"`
	}

	v := lb{
		Context:       "https://schema.org",
		Type:          []string{"LocalBusiness", "SportsActivityLocation", "EntertainmentBusiness"},
		Name:          BusinessName,
		AlternateName: "Ezy Gaming Zone",
		Description:   "Ezy Gaming is Ahmedabad's premier gaming café in Nikol, offering high-end Gaming PCs with RTX 4090, PS5 consoles on 4K displays, esports tournaments, and gaming membership plans.",
		URL:           SiteURL,
		Telephone:     Phone,
		Email:         Email,
		Image:         SiteURL + "/static/img/gaming-cafe-interior.jpg",
		Logo:          SiteURL + "/static/images/logo_without_bac.png",
		Address: address{
			Type:            "PostalAddress",
			StreetAddress:   Street,
			AddressLocality: Locality + ", " + City,
			AddressRegion:   Region,
			PostalCode:      PostalCode,
			AddressCountry:  Country,
		},
		Geo: geo{
			Type:      "GeoCoordinates",
			Latitude:  Lat,
			Longitude: Lng,
		},
		OpeningHoursSpecification: []openingHours{
			{Type: "OpeningHoursSpecification", DayOfWeek: []string{"Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"}, Opens: "00:00", Closes: "23:59"},
		},
		PriceRange: "₹₹",
		AggregateRating: aggregateRating{
			Type:        "AggregateRating",
			RatingValue: 4.8,
			ReviewCount: 127,
			BestRating:  5,
		},
		SameAs: []string{
			"https://www.instagram.com/ezygamingzone",
			"https://www.facebook.com/ezygamingzone",
			"https://twitter.com/ezygamingzone",
			"https://www.youtube.com/@ezygamingzone",
		},
		Keywords: "gaming zone ahmedabad, gaming cafe ahmedabad, gaming zone nikol, ps5 zone ahmedabad, gaming pc cafe ahmedabad, esports cafe ahmedabad, gaming arena ahmedabad",
	}
	return jsonLD(v)
}

// WebSiteSchema returns the WebSite schema with SearchAction.
func WebSiteSchema() template.HTML {
	v := map[string]any{
		"@context": "https://schema.org",
		"@type":    "WebSite",
		"name":     SiteName,
		"url":      SiteURL,
		"potentialAction": map[string]any{
			"@type":       "SearchAction",
			"target": map[string]any{
				"@type":       "EntryPoint",
				"urlTemplate": SiteURL + "/search?q={search_term_string}",
			},
			"query-input": "required name=search_term_string",
		},
	}
	return jsonLD(v)
}

// OrganizationSchema returns Organization structured data.
func OrganizationSchema() template.HTML {
	v := map[string]any{
		"@context": "https://schema.org",
		"@type":    "Organization",
		"name":     BusinessName,
		"url":      SiteURL,
		"logo":     SiteURL + "/static/images/logo_without_bac.png",
		"contactPoint": map[string]any{
			"@type":             "ContactPoint",
			"telephone":         Phone,
			"contactType":       "customer service",
			"areaServed":        "IN",
			"availableLanguage": []string{"English", "Hindi", "Gujarati"},
		},
		"address": map[string]any{
			"@type":           "PostalAddress",
			"streetAddress":   Street,
			"addressLocality": City,
			"addressRegion":   Region,
			"postalCode":      PostalCode,
			"addressCountry":  Country,
		},
		"sameAs": []string{
			"https://www.instagram.com/ezygamingzone",
			"https://www.facebook.com/ezygamingzone",
		},
	}
	return jsonLD(v)
}

// BreadcrumbSchema generates a BreadcrumbList from a slice of Breadcrumb items.
func BreadcrumbSchema(items []Breadcrumb) template.HTML {
	if len(items) == 0 {
		return ""
	}
	type listItem struct {
		Type     string `json:"@type"`
		Position int    `json:"position"`
		Name     string `json:"name"`
		Item     string `json:"item"`
	}
	elements := make([]listItem, len(items))
	for i, b := range items {
		elements[i] = listItem{
			Type:     "ListItem",
			Position: i + 1,
			Name:     b.Name,
			Item:     SiteURL + b.URL,
		}
	}
	v := map[string]any{
		"@context":        "https://schema.org",
		"@type":           "BreadcrumbList",
		"itemListElement": elements,
	}
	return jsonLD(v)
}

// FAQSchema generates FAQPage structured data from Q&A pairs.
func FAQSchema(faqs []FAQ) template.HTML {
	if len(faqs) == 0 {
		return ""
	}
	type answerObject struct {
		Type string `json:"@type"`
		Text string `json:"text"`
	}
	type question struct {
		Type           string       `json:"@type"`
		Name           string       `json:"name"`
		AcceptedAnswer answerObject `json:"acceptedAnswer"`
	}
	items := make([]question, len(faqs))
	for i, f := range faqs {
		items[i] = question{
			Type: "Question",
			Name: f.Question,
			AcceptedAnswer: answerObject{Type: "Answer", Text: f.Answer},
		}
	}
	v := map[string]any{
		"@context":   "https://schema.org",
		"@type":      "FAQPage",
		"mainEntity": items,
	}
	return jsonLD(v)
}

// FAQ is a single Q&A pair.
type FAQ struct {
	Question string
	Answer   string
}

// EventSchema generates Event structured data for a tournament.
type TournamentEvent struct {
	Name        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	Location    string
	URL         string
	Organizer   string
	EntryFee    float64
	PrizePool   float64
}

func EventSchema(e TournamentEvent) template.HTML {
	v := map[string]any{
		"@context":    "https://schema.org",
		"@type":       "Event",
		"name":        e.Name,
		"description": e.Description,
		"startDate":   e.StartDate.Format(time.RFC3339),
		"endDate":     e.EndDate.Format(time.RFC3339),
		"eventStatus": "https://schema.org/EventScheduled",
		"eventAttendanceMode": "https://schema.org/OfflineEventAttendanceMode",
		"location": map[string]any{
			"@type": "Place",
			"name":  BusinessName,
			"address": map[string]any{
				"@type":           "PostalAddress",
				"streetAddress":   Street,
				"addressLocality": City,
				"addressRegion":   Region,
				"addressCountry":  Country,
			},
		},
		"organizer": map[string]any{
			"@type": "Organization",
			"name":  e.Organizer,
			"url":   SiteURL,
		},
		"offers": map[string]any{
			"@type":         "Offer",
			"price":         fmt.Sprintf("%.0f", e.EntryFee),
			"priceCurrency": "INR",
			"availability":  "https://schema.org/InStock",
			"url":           SiteURL + "/tournaments",
		},
	}
	if e.URL != "" {
		v["url"] = e.URL
	}
	return jsonLD(v)
}

// ServiceSchema generates a Service schema for a specific offering.
type ServiceOffer struct {
	Name        string
	Description string
	URL         string
	Price       string
	Category    string
}

func ServiceSchema(s ServiceOffer) template.HTML {
	v := map[string]any{
		"@context":     "https://schema.org",
		"@type":        "Service",
		"name":         s.Name,
		"description":  s.Description,
		"url":          SiteURL + s.URL,
		"category":     s.Category,
		"provider": map[string]any{
			"@type": "LocalBusiness",
			"name":  BusinessName,
			"url":   SiteURL,
		},
		"areaServed": map[string]any{
			"@type":           "City",
			"name":            City,
			"addressRegion":   Region,
		},
		"offers": map[string]any{
			"@type":         "Offer",
			"price":         s.Price,
			"priceCurrency": "INR",
		},
	}
	return jsonLD(v)
}
