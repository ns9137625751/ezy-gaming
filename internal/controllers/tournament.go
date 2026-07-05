package controllers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/repositories"
	"github.com/nishantshekhada/ezygaming/internal/seo"
	"github.com/nishantshekhada/ezygaming/pkg/renderer"
)

type TournamentController struct {
	repo repositories.TournamentRepository
}

func NewTournamentController(repo repositories.TournamentRepository) *TournamentController {
	return &TournamentController{repo: repo}
}

func (h *TournamentController) Index(c *gin.Context) {
	all, _ := h.repo.FindAll()

	// Build Event JSON-LD for the first upcoming tournament (if any).
	var extraLDs []template.HTML
	for _, t := range all {
		if t.Status == "upcoming" {
			extraLDs = append(extraLDs, seo.EventSchema(seo.TournamentEvent{
				Name:        t.Title,
				Description: t.Description,
				StartDate:   t.StartDate,
				EndDate:     t.StartDate,
				Organizer:   seo.BusinessName,
				EntryFee:    t.EntryFee,
				PrizePool:   t.PrizePool,
				URL:         seo.SiteURL + "/tournaments",
			}))
			break
		}
	}

	faqs := []seo.FAQ{
		{Question: "Does Ezy Gaming host esports tournaments in Ahmedabad?", Answer: "Yes! Ezy Gaming Zone in Nikol, Ahmedabad hosts weekly and monthly tournaments for games like Valorant, CS2, FIFA 25, and more. Cash prizes, gaming gear, and EZY credits for winners."},
		{Question: "How do I register for a gaming tournament in Ahmedabad?", Answer: "Visit ezygaming.in/tournaments to see upcoming events, then register your team via our contact form or by visiting our gaming café in Nikol, Ahmedabad."},
		{Question: "What is the entry fee for gaming tournaments at Ezy Gaming?", Answer: "Entry fees vary by tournament — some events are free for members. Standard entry fees range from ₹50 to ₹200 per team/player."},
		{Question: "Can I host a private gaming tournament at Ezy Gaming, Ahmedabad?", Answer: "Yes! We host corporate events, college tournaments, and community events. Contact us at events@ezygaming.in or visit us in Nikol, Ahmedabad."},
	}

	extraLDs = append(extraLDs, seo.FAQSchema(faqs))

	data := seo.Build(seo.PageConfig{
		Title:       "Gaming Tournaments in Ahmedabad — Esports Events at Ezy Gaming",
		Description: "Join weekly esports tournaments at Ezy Gaming Zone, Nikol Ahmedabad. Valorant, CS2, FIFA, and more. Cash prizes, gaming gear. Register your team today.",
		Keywords:    "gaming tournament ahmedabad, esports tournament ahmedabad, gaming competition nikol, valorant tournament ahmedabad, cs2 tournament ahmedabad, gaming events gujarat",
		Path:        "/tournaments",
		Breadcrumbs: []seo.Breadcrumb{{Name: "Tournaments"}},
		ExtraJSONLDs: extraLDs,
		ExtraData:   gin.H{"Tournaments": all},
	})

	renderer.Render(c, http.StatusOK, "tournaments", data)
}
