package admin

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/middlewares"
	"github.com/nishantshekhada/ezygaming/internal/models"
	"github.com/nishantshekhada/ezygaming/internal/repositories"
)

type TournamentAdminController struct {
	repo        repositories.TournamentRepository
	contactRepo repositories.ContactRepository
}

func NewTournamentAdminController(repo repositories.TournamentRepository, contactRepo repositories.ContactRepository) *TournamentAdminController {
	return &TournamentAdminController{repo, contactRepo}
}

func (h *TournamentAdminController) Index(c *gin.Context) {
	user := middlewares.GetAdminUser(c)
	items, _ := h.repo.FindAll()
	unread, _ := h.contactRepo.CountUnread()
	renderAdmin(c, http.StatusOK, "tournament", gin.H{
		"User":           user,
		"Page":           "tournament",
		"Items":          items,
		"Flash":          c.Query("flash"),
		"UnreadContacts": unread,
	})
}

func (h *TournamentAdminController) Create(c *gin.Context) {
	t, err := parseTournamentForm(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/tournament?flash=invalid")
		return
	}
	if err := h.repo.Create(t); err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/tournament?flash=error")
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin/tournament?flash=created")
}

func (h *TournamentAdminController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/tournament")
		return
	}
	existing, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/tournament?flash=notfound")
		return
	}
	updated, err := parseTournamentForm(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/tournament?flash=invalid")
		return
	}
	existing.Title                = updated.Title
	existing.GameName             = updated.GameName
	existing.Platform             = updated.Platform
	existing.Format               = updated.Format
	existing.PrizePool            = updated.PrizePool
	existing.EntryFee             = updated.EntryFee
	existing.MaxTeams             = updated.MaxTeams
	existing.RegisteredTeams      = updated.RegisteredTeams
	existing.StartDate            = updated.StartDate
	existing.RegistrationDeadline = updated.RegistrationDeadline
	existing.Status               = updated.Status
	existing.Description          = updated.Description
	existing.Rules                = updated.Rules
	existing.IsActive             = updated.IsActive
	if err := h.repo.Update(existing); err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/tournament?flash=error")
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin/tournament?flash=updated")
}

func (h *TournamentAdminController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/tournament")
		return
	}
	h.repo.Delete(uint(id))
	c.Redirect(http.StatusSeeOther, "/admin/tournament?flash=deleted")
}

func parseTournamentForm(c *gin.Context) (*models.Tournament, error) {
	title := strings.TrimSpace(c.PostForm("title"))
	if title == "" {
		return nil, strconv.ErrSyntax
	}
	prizePool, _ := strconv.ParseFloat(c.PostForm("prize_pool"), 64)
	entryFee, _  := strconv.ParseFloat(c.PostForm("entry_fee"), 64)
	maxTeams, _  := strconv.Atoi(c.PostForm("max_teams"))
	regTeams, _  := strconv.Atoi(c.PostForm("registered_teams"))

	var startDate, regDeadline time.Time
	if v := c.PostForm("start_date"); v != "" {
		startDate, _ = time.ParseInLocation("2006-01-02", v, time.Local)
	}
	if v := c.PostForm("registration_deadline"); v != "" {
		regDeadline, _ = time.ParseInLocation("2006-01-02", v, time.Local)
	}

	status := c.PostForm("status")
	if status == "" {
		status = "upcoming"
	}

	return &models.Tournament{
		Title:                title,
		GameName:             strings.TrimSpace(c.PostForm("game_name")),
		Platform:             c.PostForm("platform"),
		Format:               strings.TrimSpace(c.PostForm("format")),
		PrizePool:            prizePool,
		EntryFee:             entryFee,
		MaxTeams:             maxTeams,
		RegisteredTeams:      regTeams,
		StartDate:            startDate,
		RegistrationDeadline: regDeadline,
		Status:               status,
		Description:          strings.TrimSpace(c.PostForm("description")),
		Rules:                strings.TrimSpace(c.PostForm("rules")),
		IsActive:             c.PostForm("is_active") == "1",
	}, nil
}
