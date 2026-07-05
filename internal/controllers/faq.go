package controllers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/models"
	"github.com/nishantshekhada/ezygaming/internal/repositories"
	"github.com/nishantshekhada/ezygaming/internal/seo"
	"github.com/nishantshekhada/ezygaming/pkg/renderer"
)

type FAQController struct {
	repo repositories.FAQRepository
}

func NewFAQController(repo repositories.FAQRepository) *FAQController {
	return &FAQController{repo: repo}
}

func (h *FAQController) Index(c *gin.Context) {
	faqs, _ := h.repo.FindActive()

	var faqLDs []seo.FAQ
	for _, f := range faqs {
		faqLDs = append(faqLDs, seo.FAQ{Question: f.Question, Answer: f.Answer})
	}

	// Group by category preserving first-seen order
	catOrder := []string{}
	grouped := map[string][]models.FAQ{}
	for _, f := range faqs {
		cat := f.Category
		if cat == "" {
			cat = "General"
		}
		if _, ok := grouped[cat]; !ok {
			catOrder = append(catOrder, cat)
		}
		grouped[cat] = append(grouped[cat], f)
	}

	var extraLDs []template.HTML
	if len(faqLDs) > 0 {
		extraLDs = append(extraLDs, seo.FAQSchema(faqLDs))
	}

	data := seo.Build(seo.PageConfig{
		Title:        "FAQs — Ezy Gaming Zone Ahmedabad",
		Description:  "Frequently asked questions about Ezy Gaming Zone, Nikol Ahmedabad — pricing, bookings, membership, tournaments, and more.",
		Keywords:     "gaming cafe faq ahmedabad, ezy gaming questions, gaming zone help nikol",
		Path:         "/faqs",
		Breadcrumbs:  []seo.Breadcrumb{{Name: "FAQs"}},
		ExtraJSONLDs: extraLDs,
		ExtraData: gin.H{
			"FAQs":     faqs,
			"Grouped":  grouped,
			"CatOrder": catOrder,
		},
	})

	renderer.Render(c, http.StatusOK, "faqs", data)
}
