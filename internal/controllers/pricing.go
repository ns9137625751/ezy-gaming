package controllers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/seo"
	"github.com/nishantshekhada/ezygaming/pkg/renderer"
)

type PricingController struct{}

func NewPricingController() *PricingController { return &PricingController{} }

type PricingTier struct {
	Name        string
	Description string
	HourlyRate  int
	Features    []string
	PriceColor  string
	IconBg      string
	IconColor   string
	IsPopular   bool
}

func (h *PricingController) Index(c *gin.Context) {
	pcTiers := []PricingTier{
		{
			Name: "Standard", HourlyRate: 60, PriceColor: "#06b6d4",
			Description: "Great for casual gaming sessions",
			IconBg: "rgba(6,182,212,.15)", IconColor: "#06b6d4",
			Features: []string{"RTX 3090 GPU", "Intel i7-12700K CPU", "16GB DDR4 RAM", "24\" 144Hz IPS Display", "1TB SSD Storage", "Gaming Peripherals Included"},
		},
		{
			Name: "Pro", HourlyRate: 70, PriceColor: "#8b5cf6", IsPopular: true,
			Description: "For competitive and enthusiast gamers",
			IconBg: "rgba(139,92,246,.15)", IconColor: "#8b5cf6",
			Features: []string{"RTX 4080 GPU", "Intel i7-13700K CPU", "32GB DDR5 RAM", "27\" 165Hz IPS Display", "1TB NVMe SSD", "HyperX RGB Peripherals"},
		},
		{
			Name: "Ultra", HourlyRate: 80, PriceColor: "#f59e0b",
			Description: "Maximum performance for elite players",
			IconBg: "rgba(245,158,11,.15)", IconColor: "#f59e0b",
			Features: []string{"RTX 4090 GPU", "Intel i9-13900K CPU", "32GB DDR5 RAM", "27\" 240Hz IPS Display", "2TB NVMe SSD", "HyperX RGB Peripherals"},
		},
	}

	ps5Tiers := []PricingTier{
		{
			Name: "Standard Console", HourlyRate: 50, PriceColor: "#8b5cf6",
			Description: "Next-gen gaming on a massive screen",
			IconBg: "rgba(139,92,246,.15)", IconColor: "#8b5cf6",
			Features: []string{"PlayStation 5 Console", "50\" 4K QLED Display", "2× DualSense Controllers", "Up to 4 Players", "PlayStation Plus Access", "100+ PS5 Titles"},
		},
		{
			Name: "Premium Console", HourlyRate: 60, PriceColor: "#06b6d4", IsPopular: true,
			Description: "The ultimate PS5 experience",
			IconBg: "rgba(6,182,212,.15)", IconColor: "#06b6d4",
			Features: []string{"PlayStation 5 Console", "55\" 4K OLED Display", "2× DualSense Controllers", "Dolby Atmos 3D Audio", "Up to 4 Players", "100+ PS5 Titles"},
		},
	}

	faqs := []seo.FAQ{
		{Question: "How much does it cost to play at a gaming café in Ahmedabad?", Answer: "At Ezy Gaming Zone in Nikol, Ahmedabad, Gaming PC sessions start at ₹60/hr (Standard), ₹70/hr (Pro), and ₹80/hr (Ultra RTX 4090). PS5 pods start at ₹50/hr."},
		{Question: "Is there a minimum session time at Ezy Gaming?", Answer: "Minimum session is 1 hour. After that you can extend in 30-minute increments at the same hourly rate."},
		{Question: "Do gaming café prices in Ahmedabad include peripherals?", Answer: "Yes! All peripherals — mechanical keyboard, gaming mouse, and headset — are included in every session at no extra cost."},
		{Question: "Are there membership plans at Ezy Gaming Zone Ahmedabad?", Answer: "Yes. Membership plans start at ₹499/month (Bronze) up to ₹2499/month (Platinum) with unlimited hours. Members get 10–30% off all sessions plus tournament access."},
	}

	data := seo.Build(seo.PageConfig{
		Title:       "Gaming Café Prices in Ahmedabad — PC & PS5 Rates",
		Description: "Gaming PC from ₹60/hr, PS5 from ₹50/hr at Ezy Gaming Zone, Nikol Ahmedabad. Transparent hourly rates, no hidden fees. Monthly memberships from ₹499.",
		Keywords:    "gaming cafe price ahmedabad, gaming zone rates ahmedabad, gaming pc hourly rate ahmedabad, ps5 cafe price ahmedabad, gaming membership ahmedabad",
		Path:        "/pricing",
		Breadcrumbs: []seo.Breadcrumb{{Name: "Pricing"}},
		ExtraJSONLDs: []template.HTML{seo.FAQSchema(faqs)},
		ExtraData: gin.H{
			"PCTiers":  pcTiers,
			"PS5Tiers": ps5Tiers,
		},
	})

	renderer.Render(c, http.StatusOK, "pricing", data)
}
