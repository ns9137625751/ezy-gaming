package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/models"
	"github.com/nishantshekhada/ezygaming/internal/repositories"
	"html/template"
	"github.com/nishantshekhada/ezygaming/internal/seo"
	"github.com/nishantshekhada/ezygaming/pkg/renderer"
)

type HomeController struct {
	pcRepo         repositories.GamingPCRepository
	tournamentRepo repositories.TournamentRepository
	membershipRepo repositories.MembershipPlanRepository
	galleryRepo    repositories.GalleryRepository
}

func NewHomeController(
	pcRepo repositories.GamingPCRepository,
	tRepo repositories.TournamentRepository,
	mRepo repositories.MembershipPlanRepository,
	gRepo repositories.GalleryRepository,
) *HomeController {
	return &HomeController{
		pcRepo:         pcRepo,
		tournamentRepo: tRepo,
		membershipRepo: mRepo,
		galleryRepo:    gRepo,
	}
}

func (h *HomeController) Index(c *gin.Context) {
	pcs, _ := h.pcRepo.FindActive()
	tournaments, _ := h.tournamentRepo.FindUpcoming(3)
	memberships, _ := h.membershipRepo.FindActive()
	gallery, _ := h.galleryRepo.FindFeatured(6)

	faqs := []seo.FAQ{
		{Question: "Where is Ezy Gaming located?", Answer: "Ezy Gaming is located at Vishala Supreme, 5-6, Sardar Patel Ring Rd, near Torrent Power subStation, Nikol, Ahmedabad, Gujarat 380049 — easily accessible via the Nikol BRTS stop."},
		{Question: "What gaming equipment does Ezy Gaming have?", Answer: "We have 20 high-end Gaming PCs with RTX 4090, Intel i9-13900K, 32GB DDR5 RAM and 240Hz displays, plus 20 PS5 consoles connected to 4K OLED and QLED TVs up to 55\"."},
		{Question: "What are the hourly rates at Ezy Gaming?", Answer: "Gaming PC rates start at ₹60/hr for Standard PCs, ₹70/hr for Pro PCs, and ₹80/hr for Ultra PCs. PS5 pods start at ₹50/hr. We also offer monthly membership plans from ₹499."},
		{Question: "Can I book a gaming station online?", Answer: "Yes! You can book any Gaming PC or PS5 pod online at ezygaming.in/book. Pay at the venue when you arrive. Free cancellation up to 2 hours before your session."},
		{Question: "Does Ezy Gaming host tournaments?", Answer: "Yes, we host weekly and monthly esports tournaments for games like Valorant, CS2, FIFA, and more. Cash prizes, gaming gear, and EZY credit for top finishers."},
	}

	data := seo.Build(seo.PageConfig{
		Title:       "Gaming Zone in Nikol, Ahmedabad",
		Description: "Ezy Gaming Zone — Ahmedabad's best gaming café in Nikol. 20 RTX 4090 Gaming PCs, 20 PS5 consoles on 55\" 4K OLED, weekly esports tournaments. Book online now.",
		Keywords:    "gaming zone ahmedabad, gaming cafe ahmedabad, gaming zone nikol, ps5 zone ahmedabad, esports cafe ahmedabad, gaming pc cafe nikol, gaming arena ahmedabad",
		Path:        "/",
		OGType:      "website",
		ExtraJSONLDs: []template.HTML{seo.FAQSchema(faqs), seo.OrganizationSchema()},
		ExtraData: gin.H{
			"PCs":          pcs,
			"Tournaments":  tournaments,
			"Memberships":  memberships,
			"Gallery":      gallery,
			"PCCount":      countActive(pcs),
			"Testimonials": defaultTestimonials(),
			"PopularGames": defaultGames(),
			"HomeFAQs":     faqs,
		},
	})

	renderer.Render(c, http.StatusOK, "home", data)
}

func countActive(pcs []models.GamingPC) int {
	n := 0
	for _, pc := range pcs {
		if pc.Status == "available" {
			n++
		}
	}
	return n
}

type Testimonial struct {
	Name   string
	Role   string
	Text   string
	Rating int
}

type PopularGame struct {
	Name     string
	Platform string
	Icon     string
}

func defaultTestimonials() []Testimonial {
	return []Testimonial{
		{Name: "Arjun Patel", Role: "Pro Gamer, Ahmedabad", Text: "Best gaming café in Nikol! The RTX 4090 setup is insane — zero lag, buttery smooth performance. I come here every weekend for Valorant ranked matches.", Rating: 5},
		{Name: "Priya Shah", Role: "College Student, Gujarat", Text: "Love the PS5 zone! The 55\" OLED TVs make FIFA look absolutely stunning. Super clean setup and great staff. Best gaming experience in Ahmedabad!", Rating: 5},
		{Name: "Karan Mehta", Role: "Streamer, Nikol", Text: "The internet speed and hardware here are top-notch. I've streamed multiple sessions from Ezy Gaming with zero issues. Best gaming lounge in Ahmedabad by far.", Rating: 5},
	}
}

func defaultGames() []PopularGame {
	return []PopularGame{
		{Name: "Valorant", Platform: "PC", Icon: "🎯"},
		{Name: "CS2", Platform: "PC", Icon: "💣"},
		{Name: "GTA V", Platform: "PC", Icon: "🚗"},
		{Name: "FIFA 25", Platform: "PS5", Icon: "⚽"},
		{Name: "Call of Duty", Platform: "Both", Icon: "🔫"},
		{Name: "Spider-Man 2", Platform: "PS5", Icon: "🕷️"},
		{Name: "God of War", Platform: "PS5", Icon: "⚔️"},
		{Name: "Minecraft", Platform: "PC", Icon: "🧱"},
	}
}
