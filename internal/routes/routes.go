package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	adminCtrl "github.com/nishantshekhada/ezygaming/internal/controllers/admin"
	customerCtrl "github.com/nishantshekhada/ezygaming/internal/controllers/customer"
	"github.com/nishantshekhada/ezygaming/internal/controllers"
	"github.com/nishantshekhada/ezygaming/internal/middlewares"
	"github.com/nishantshekhada/ezygaming/internal/repositories"
	"gorm.io/gorm"
)

func Setup(r *gin.Engine, db *gorm.DB, jwtSecret string) {
	// Global middleware
	r.Use(middlewares.SEOHeaders())
	r.Use(middlewares.CustomerSessionInfo(jwtSecret))

	// Repositories
	pcRepo      := repositories.NewGamingPCRepository(db)
	ps5Repo     := repositories.NewPS5Repository(db)
	tRepo       := repositories.NewTournamentRepository(db)
	mRepo       := repositories.NewMembershipPlanRepository(db)
	gRepo       := repositories.NewGalleryRepository(db)
	contactRepo := repositories.NewContactRepository(db)
	slotRepo    := repositories.NewTimeSlotRepository(db)
	bookingRepo := repositories.NewBookingRepository(db)
	adminRepo    := repositories.NewAdminUserRepository(db)
	customerRepo := repositories.NewCustomerRepository(db)
	inventoryRepo := repositories.NewInventoryRepository(db)
	faqRepo       := repositories.NewFAQRepository(db)

	// Controllers — public site
	homeCtrl       := controllers.NewHomeController(pcRepo, tRepo, mRepo, gRepo)
	pcCtrl         := controllers.NewGamingPCController(pcRepo)
	ps5Ctrl        := controllers.NewPS5Controller(ps5Repo)
	pricingCtrl    := controllers.NewPricingController()
	membershipCtrl := controllers.NewMembershipController(mRepo)
	tournamentCtrl := controllers.NewTournamentController(tRepo)
	contactCtrl    := controllers.NewContactController(contactRepo)
	bookingCtrl    := controllers.NewBookingController(inventoryRepo, slotRepo, bookingRepo)
	seoCtrl        := controllers.NewSEOController()
	faqCtrl        := controllers.NewFAQController(faqRepo)
	galleryCtrl    := controllers.NewGalleryController(gRepo)
	pagesCtrl      := controllers.NewPagesController()

	// Controllers — admin
	authCtrl  := adminCtrl.NewAuthController(adminRepo, jwtSecret)
	dashCtrl  := adminCtrl.NewDashboardController(bookingRepo, contactRepo, inventoryRepo)

	// Controllers — customer portal
	custAuthCtrl    := customerCtrl.NewAuthController(customerRepo, jwtSecret)
	custProfileCtrl := customerCtrl.NewProfileController(customerRepo, bookingRepo)
	inventoryCtrl     := adminCtrl.NewInventoryController(inventoryRepo, contactRepo)
	tournamentAdminCtrl := adminCtrl.NewTournamentAdminController(tRepo, contactRepo)

	// ── SEO & crawl ────────────────────────────────────────────────
	r.GET("/sitemap.xml", seoCtrl.Sitemap)
	r.GET("/robots.txt", seoCtrl.RobotsTxt)
	r.GET("/humans.txt", seoCtrl.HumansTxt)
	r.GET("/.well-known/security.txt", seoCtrl.SecurityTxt)

	// ── Public pages ────────────────────────────────────────────────
	r.GET("/", homeCtrl.Index)
	r.GET("/gaming-pcs", pcCtrl.Index)
	r.GET("/ps5-zone", ps5Ctrl.Index)
	r.GET("/pricing", pricingCtrl.Index)
	r.GET("/memberships", membershipCtrl.Index)
	r.GET("/tournaments", tournamentCtrl.Index)
	r.GET("/faqs", faqCtrl.Index)
	r.GET("/gallery", galleryCtrl.Index)
	r.GET("/about", pagesCtrl.About)
	r.GET("/privacy", pagesCtrl.Privacy)
	r.GET("/terms", pagesCtrl.Terms)
	r.GET("/contact", contactCtrl.Index)
	r.POST("/contact", contactCtrl.Submit)
	r.GET("/book", bookingCtrl.Index)
	r.POST("/book", bookingCtrl.Submit)

	// ── Admin panel ─────────────────────────────────────────────────
	adm := r.Group("/admin")
	adm.GET("/login", authCtrl.LoginPage)
	adm.POST("/login", authCtrl.Login)

	// Authenticated admin routes
	auth := adm.Group("", middlewares.AdminAuth(jwtSecret))
	auth.GET("", func(c *gin.Context) { c.Redirect(http.StatusSeeOther, "/admin/dashboard") })
	auth.GET("/logout", authCtrl.Logout)
	auth.GET("/dashboard", dashCtrl.Dashboard)

	// Bookings — admin + gamer
	auth.GET("/bookings", dashCtrl.Bookings)
	auth.POST("/bookings/new", dashCtrl.CreateBooking)
	auth.POST("/bookings/:id/status", middlewares.RequireRole("admin"), dashCtrl.UpdateBookingStatus)

	// Inventory — admin only
	auth.GET("/inventory", middlewares.RequireRole("admin"), inventoryCtrl.Index)
	auth.POST("/inventory", middlewares.RequireRole("admin"), inventoryCtrl.Create)
	auth.POST("/inventory/:id/update", middlewares.RequireRole("admin"), inventoryCtrl.Update)
	auth.POST("/inventory/:id/delete", middlewares.RequireRole("admin"), inventoryCtrl.Delete)

	// Tournament — admin only
	auth.GET("/tournament", middlewares.RequireRole("admin"), tournamentAdminCtrl.Index)
	auth.POST("/tournament", middlewares.RequireRole("admin"), tournamentAdminCtrl.Create)
	auth.POST("/tournament/:id/update", middlewares.RequireRole("admin"), tournamentAdminCtrl.Update)
	auth.POST("/tournament/:id/delete", middlewares.RequireRole("admin"), tournamentAdminCtrl.Delete)

	// Contacts — admin only
	auth.GET("/contacts", middlewares.RequireRole("admin"), dashCtrl.Contacts)
	auth.POST("/contacts/:id/read", middlewares.RequireRole("admin"), dashCtrl.MarkContactRead)

	// ── Customer portal ─────────────────────────────────────────────
	r.GET("/login", custAuthCtrl.LoginPage)
	r.POST("/login", custAuthCtrl.Login)
	r.GET("/register", custAuthCtrl.RegisterPage)
	r.POST("/register", custAuthCtrl.Register)
	r.GET("/logout", custAuthCtrl.Logout)

	cust := r.Group("/me", middlewares.CustomerAuth(jwtSecret))
	cust.GET("", custProfileCtrl.Me)
}
