package controllers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/repositories"
	"github.com/nishantshekhada/ezygaming/internal/seo"
	"github.com/nishantshekhada/ezygaming/pkg/renderer"
)

type MembershipController struct {
	repo repositories.MembershipPlanRepository
}

func NewMembershipController(repo repositories.MembershipPlanRepository) *MembershipController {
	return &MembershipController{repo: repo}
}

func (h *MembershipController) Index(c *gin.Context) {
	plans, _ := h.repo.FindActive()

	faqs := []seo.FAQ{
		{Question: "Does Ezy Gaming Zone offer gaming memberships in Ahmedabad?", Answer: "Yes! We offer Bronze (₹499/mo), Silver (₹799/mo), Gold (₹1499/mo), and Platinum (₹2499/mo) membership plans for our gaming café in Nikol, Ahmedabad."},
		{Question: "What benefits do gaming café membership plans include?", Answer: "Members get included gaming hours, 10–30% off hourly rates, priority booking, tournament access, free snack vouchers, and for Platinum members — dedicated VIP stations."},
		{Question: "How do I activate my Ezy Gaming membership in Ahmedabad?", Answer: "Visit our gaming café at Nikol, Ahmedabad or contact us at hello@ezygaming.in. Memberships activate instantly and can be cancelled anytime with no lock-in."},
		{Question: "Can I use my membership for both PC and PS5 sessions?", Answer: "Yes! Gold and Platinum plans include both PC and PS5 session hours. Bronze and Silver plans cover PC sessions with discounts applicable to PS5 rentals."},
	}

	data := seo.Build(seo.PageConfig{
		Title:       "Gaming Membership Plans in Ahmedabad — Ezy Gaming Zone",
		Description: "Save more with Ezy Gaming membership plans in Nikol, Ahmedabad. Bronze ₹499/mo to Platinum ₹2499/mo. Unlimited hours, 30% off, priority booking, tournament access.",
		Keywords:    "gaming membership ahmedabad, gaming cafe membership nikol, gaming subscription ahmedabad, gaming pass ahmedabad, gaming zone membership gujarat",
		Path:        "/memberships",
		Breadcrumbs: []seo.Breadcrumb{{Name: "Memberships"}},
		ExtraJSONLDs: []template.HTML{seo.FAQSchema(faqs)},
		ExtraData:   gin.H{"Plans": plans},
	})

	renderer.Render(c, http.StatusOK, "membership", data)
}
