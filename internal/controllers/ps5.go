package controllers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/repositories"
	"github.com/nishantshekhada/ezygaming/internal/seo"
	"github.com/nishantshekhada/ezygaming/pkg/renderer"
)

type PS5Controller struct {
	ps5Repo repositories.PS5Repository
}

func NewPS5Controller(ps5Repo repositories.PS5Repository) *PS5Controller {
	return &PS5Controller{ps5Repo: ps5Repo}
}

func (h *PS5Controller) Index(c *gin.Context) {
	consoles, _ := h.ps5Repo.FindActive()

	available := 0
	for _, con := range consoles {
		if con.Status == "available" {
			available++
		}
	}

	faqs := []seo.FAQ{
		{Question: "Does Ezy Gaming have a PS5 zone in Ahmedabad?", Answer: "Yes! Ezy Gaming in Nikol, Ahmedabad has 20 PS5 consoles. Premium pods have 55\" 4K OLED TVs at ₹60/hr, and Standard pods have 50\" 4K QLED TVs at ₹50/hr."},
		{Question: "Which PS5 games are available at Ezy Gaming?", Answer: "We have 100+ PS5 titles including God of War Ragnarök, Spider-Man 2, FIFA 25, Gran Turismo 7, Elden Ring, Final Fantasy XVI, and many more through PlayStation Plus."},
		{Question: "Can I play PS5 multiplayer at Ezy Gaming?", Answer: "Absolutely! Each PS5 pod includes 2 DualSense controllers, so you can enjoy local multiplayer with a friend. Our layout also supports cross-pod gaming."},
		{Question: "Where is the PS5 zone in Nikol, Ahmedabad?", Answer: "Our PS5 zone is at Ezy Gaming, Vishala Supreme, 5-6, Sardar Patel Ring Rd, near Torrent Power subStation, Nikol, Ahmedabad 380049. Nikol BRTS stop nearby. Parking available."},
	}

	ps5Service := seo.ServiceSchema(seo.ServiceOffer{
		Name:        "PS5 Gaming Zone in Ahmedabad",
		Description: "PlayStation 5 gaming pods with 4K OLED and QLED displays at Ezy Gaming Zone, Nikol, Ahmedabad. Hourly rates from ₹50.",
		URL:         "/ps5-zone",
		Price:       "50",
		Category:    "Entertainment",
	})

	data := seo.Build(seo.PageConfig{
		Title:       "PS5 Zone in Nikol, Ahmedabad — 4K OLED Gaming",
		Description: "PS5 gaming zone at Ezy Gaming, Nikol Ahmedabad. 20 PS5 consoles on 55\" 4K OLED & 50\" QLED displays. 100+ titles, DualSense controllers. From ₹50/hr.",
		Keywords:    "ps5 zone ahmedabad, ps5 gaming cafe ahmedabad, ps5 zone nikol, ps5 cafe gujarat, playstation 5 ahmedabad, ps5 gaming lounge ahmedabad",
		Path:        "/ps5-zone",
		Breadcrumbs: []seo.Breadcrumb{
			{Name: "PS5 Zone"},
		},
		ExtraJSONLDs: []template.HTML{seo.FAQSchema(faqs), ps5Service},
		ExtraData: gin.H{
			"Consoles":  consoles,
			"Available": available,
			"Total":     len(consoles),
		},
	})

	renderer.Render(c, http.StatusOK, "ps5-zone", data)
}
