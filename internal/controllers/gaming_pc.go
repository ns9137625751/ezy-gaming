package controllers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/repositories"
	"github.com/nishantshekhada/ezygaming/internal/seo"
	"github.com/nishantshekhada/ezygaming/pkg/renderer"
)

type GamingPCController struct {
	pcRepo repositories.GamingPCRepository
}

func NewGamingPCController(pcRepo repositories.GamingPCRepository) *GamingPCController {
	return &GamingPCController{pcRepo: pcRepo}
}

func (h *GamingPCController) Index(c *gin.Context) {
	pcs, _ := h.pcRepo.FindActive()

	available, occupied, maintenance := 0, 0, 0
	for _, pc := range pcs {
		switch pc.Status {
		case "available":
			available++
		case "occupied":
			occupied++
		default:
			maintenance++
		}
	}

	faqs := []seo.FAQ{
		{Question: "What Gaming PC specs are available at Ezy Gaming in Ahmedabad?", Answer: "We offer three PC tiers: Ultra (RTX 4090, i9-13900K, 32GB DDR5, 240Hz display) at ₹80/hr, Pro (RTX 4080, i7-13700K) at ₹70/hr, and Standard (RTX 3090, i7-12700K, 144Hz) at ₹60/hr."},
		{Question: "How many Gaming PCs does Ezy Gaming have?", Answer: "Ezy Gaming Zone has 20 high-end Gaming PCs across three performance tiers — 8 Ultra, 6 Pro, and 6 Standard stations."},
		{Question: "Can I book a Gaming PC in Nikol, Ahmedabad online?", Answer: "Yes! Book any available PC station at ezygaming.in/book. Pay at the venue. Minimum 1-hour session with 30-minute extensions available."},
		{Question: "Are gaming peripherals included with PC rentals?", Answer: "Yes — mechanical keyboard, gaming mouse (Logitech/Razer), and gaming headset are included at no extra charge with every PC session."},
	}

	pcService := seo.ServiceSchema(seo.ServiceOffer{
		Name:        "Gaming PC Rental in Ahmedabad",
		Description: "Rent high-end Gaming PCs with RTX 4090, i9-13900K, 32GB DDR5 RAM at Ezy Gaming Zone in Nikol, Ahmedabad. Hourly rates from ₹60.",
		URL:         "/gaming-pcs",
		Price:       "60",
		Category:    "Entertainment",
	})

	data := seo.Build(seo.PageConfig{
		Title:       "Gaming PCs in Nikol, Ahmedabad — RTX 4090 Stations",
		Description: "Rent high-end Gaming PCs at Ezy Gaming Zone, Nikol Ahmedabad. 20 stations: RTX 4090 Ultra (₹80/hr), RTX 4080 Pro (₹70/hr), RTX 3090 Standard (₹60/hr). Book online.",
		Keywords:    "gaming pc cafe ahmedabad, gaming pc zone nikol, gaming pc rental ahmedabad, RTX 4090 gaming cafe, esports PC ahmedabad, gaming PC nikol ahmedabad",
		Path:        "/gaming-pcs",
		Breadcrumbs: []seo.Breadcrumb{
			{Name: "Gaming PCs"},
		},
		ExtraJSONLDs: []template.HTML{seo.FAQSchema(faqs), pcService},
		ExtraData: gin.H{
			"PCs":         pcs,
			"Available":   available,
			"Occupied":    occupied,
			"Maintenance": maintenance,
			"Total":       len(pcs),
		},
	})

	renderer.Render(c, http.StatusOK, "gaming-pcs", data)
}
