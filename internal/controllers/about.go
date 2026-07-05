package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/seo"
	"github.com/nishantshekhada/ezygaming/pkg/renderer"
)

type AboutController struct{}

func NewAboutController() *AboutController { return &AboutController{} }

type TeamMember struct {
	Name  string
	Role  string
	Emoji string
}

func (h *AboutController) Index(c *gin.Context) {
	team := []TeamMember{
		{Name: "Rahul Sharma", Role: "Founder & CEO", Emoji: "🎮"},
		{Name: "Priya Patel", Role: "Operations Manager", Emoji: "⚡"},
		{Name: "Arjun Mehta", Role: "Head of Esports", Emoji: "🏆"},
		{Name: "Neha Joshi", Role: "Community Manager", Emoji: "🌟"},
	}

	milestones := []gin.H{
		{"year": "2023", "title": "Founded", "desc": "Ezy Gaming opened its doors in Nikol, Ahmedabad with 10 gaming PCs and a dream."},
		{"year": "2024", "title": "PS5 Zone Launch", "desc": "Expanded to 20 PS5 consoles on 4K displays, becoming Ahmedabad's largest console zone."},
		{"year": "2024", "title": "First Tournament", "desc": "Hosted our first esports tournament with 64 teams competing for ₹50,000 in prizes."},
		{"year": "2025", "title": "RTX 4090 Upgrade", "desc": "Upgraded all 20 Gaming PCs to RTX 4090 — the most powerful gaming setup in Ahmedabad."},
	}

	faqs := []seo.FAQ{
		{Question: "Who founded Ezy Gaming?", Answer: "Ezy Gaming was founded by gamers, for gamers, in Nikol, Ahmedabad. Our team is passionate about bringing premium gaming experiences to the community."},
		{Question: "What is Ezy Gaming's mission?", Answer: "Our mission is to give every gamer in Ahmedabad access to top-tier hardware — RTX 4090 PCs, PS5 consoles on 4K displays — in a welcoming, high-energy environment."},
		{Question: "Where is Ezy Gaming located?", Answer: "Ezy Gaming is at Vishala Supreme, 5-6, Sardar Patel Ring Rd, near Torrent Power subStation, Nikol, Ahmedabad, Gujarat 380049."},
	}

	data := seo.Build(seo.PageConfig{
		Title:       "About Ezy Gaming — Ahmedabad's Premier Gaming Café",
		Description: "Learn about Ezy Gaming, Ahmedabad's premier gaming café in Nikol. Our story, team, and mission to bring RTX 4090 PCs and PS5 consoles to every gamer in Gujarat.",
		Keywords:    "about ezy gaming, gaming cafe ahmedabad story, gaming zone nikol team, ezy gaming history",
		Path:        "/about",
		Breadcrumbs: []seo.Breadcrumb{{Name: "About"}},
		ExtraData: gin.H{
			"Team":       team,
			"Milestones": milestones,
			"FAQs":       faqs,
		},
	})

	renderer.Render(c, http.StatusOK, "about", data)
}
