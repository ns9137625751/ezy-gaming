package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/auth"
	"github.com/nishantshekhada/ezygaming/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	repo   repositories.AdminUserRepository
	secret string
}

func NewAuthController(repo repositories.AdminUserRepository, secret string) *AuthController {
	return &AuthController{repo: repo, secret: secret}
}

func (h *AuthController) LoginPage(c *gin.Context) {
	// Already logged in → redirect
	if tok, err := c.Cookie(auth.CookieName); err == nil && tok != "" {
		if _, err := auth.Verify(tok, h.secret); err == nil {
			c.Redirect(http.StatusSeeOther, "/admin/dashboard")
			return
		}
	}
	renderLogin(c, http.StatusOK, gin.H{"Error": ""})
}

func (h *AuthController) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if email == "" || password == "" {
		renderLogin(c, http.StatusUnprocessableEntity, gin.H{"Error": "Email and password are required."})
		return
	}

	user, err := h.repo.FindByEmail(email)
	if err != nil {
		renderLogin(c, http.StatusUnauthorized, gin.H{"Error": "Invalid email or password."})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		renderLogin(c, http.StatusUnauthorized, gin.H{"Error": "Invalid email or password."})
		return
	}

	token, err := auth.Sign(auth.SessionUser{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, h.secret)
	if err != nil {
		renderLogin(c, http.StatusInternalServerError, gin.H{"Error": "Could not create session. Try again."})
		return
	}

	// HttpOnly cookie, 8-hour session
	c.SetCookie(auth.CookieName, token, 60*60*8, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/admin/dashboard")
}

func (h *AuthController) Logout(c *gin.Context) {
	c.SetCookie(auth.CookieName, "", -1, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/admin/login")
}
