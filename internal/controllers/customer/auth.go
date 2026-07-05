package customer

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/auth"
	"github.com/nishantshekhada/ezygaming/internal/models"
	"github.com/nishantshekhada/ezygaming/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	repo   repositories.CustomerRepository
	secret string
}

func NewAuthController(repo repositories.CustomerRepository, secret string) *AuthController {
	return &AuthController{repo: repo, secret: secret}
}

func (h *AuthController) LoginPage(c *gin.Context) {
	if tok, err := c.Cookie(auth.CookieName); err == nil && tok != "" {
		if _, err := auth.Verify(tok, h.secret); err == nil {
			c.Redirect(http.StatusSeeOther, "/admin/dashboard")
			return
		}
	}
	if tok, err := c.Cookie(auth.CustomerCookieName); err == nil && tok != "" {
		if _, err := auth.Verify(tok, h.secret); err == nil {
			c.Redirect(http.StatusSeeOther, "/me")
			return
		}
	}
	renderAuth(c, http.StatusOK, "login", gin.H{"Error": "", "Tab": "login"})
}

func (h *AuthController) Login(c *gin.Context) {
	email := strings.TrimSpace(c.PostForm("email"))
	password := c.PostForm("password")

	if email == "" || password == "" {
		renderAuth(c, http.StatusUnprocessableEntity, "login", gin.H{"Error": "Email and password are required.", "Tab": "login"})
		return
	}

	cust, err := h.repo.FindByEmail(email)
	if err != nil {
		renderAuth(c, http.StatusUnauthorized, "login", gin.H{"Error": "Invalid email or password.", "Tab": "login"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(cust.Password), []byte(password)); err != nil {
		renderAuth(c, http.StatusUnauthorized, "login", gin.H{"Error": "Invalid email or password.", "Tab": "login"})
		return
	}

	h.setSessionAndRedirect(c, cust)
}

func (h *AuthController) RegisterPage(c *gin.Context) {
	if tok, err := c.Cookie(auth.CookieName); err == nil && tok != "" {
		if _, err := auth.Verify(tok, h.secret); err == nil {
			c.Redirect(http.StatusSeeOther, "/admin/dashboard")
			return
		}
	}
	if tok, err := c.Cookie(auth.CustomerCookieName); err == nil && tok != "" {
		if _, err := auth.Verify(tok, h.secret); err == nil {
			c.Redirect(http.StatusSeeOther, "/me")
			return
		}
	}
	renderAuth(c, http.StatusOK, "login", gin.H{"Error": "", "Tab": "register"})
}

func (h *AuthController) Register(c *gin.Context) {
	name := strings.TrimSpace(c.PostForm("name"))
	email := strings.TrimSpace(c.PostForm("email"))
	phone := strings.TrimSpace(c.PostForm("phone"))
	password := c.PostForm("password")

	if name == "" || email == "" || password == "" {
		renderAuth(c, http.StatusUnprocessableEntity, "login", gin.H{"Error": "All fields are required.", "Tab": "register"})
		return
	}
	if len(password) < 6 {
		renderAuth(c, http.StatusUnprocessableEntity, "login", gin.H{"Error": "Password must be at least 6 characters.", "Tab": "register"})
		return
	}

	// Backend: phone must be exactly 10 digits
	if len(phone) != 10 {
		renderAuth(c, http.StatusUnprocessableEntity, "login", gin.H{"Error": "Mobile number must be exactly 10 digits.", "Tab": "register"})
		return
	}
	for _, ch := range phone {
		if ch < '0' || ch > '9' {
			renderAuth(c, http.StatusUnprocessableEntity, "login", gin.H{"Error": "Mobile number must contain only digits.", "Tab": "register"})
			return
		}
	}

	if _, err := h.repo.FindByEmail(email); err == nil {
		renderAuth(c, http.StatusConflict, "login", gin.H{"Error": "An account with this email already exists.", "Tab": "register"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		renderAuth(c, http.StatusInternalServerError, "login", gin.H{"Error": "Could not create account. Try again.", "Tab": "register"})
		return
	}

	newCust := &models.Customer{Name: name, Email: email, Phone: phone, Password: string(hash)}
	if err := h.repo.Create(newCust); err != nil {
		renderAuth(c, http.StatusInternalServerError, "login", gin.H{"Error": "Could not create account. Try again.", "Tab": "register"})
		return
	}

	h.setSessionAndRedirect(c, newCust)
}

func (h *AuthController) Logout(c *gin.Context) {
	c.SetCookie(auth.CustomerCookieName, "", -1, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/login")
}

func (h *AuthController) setSessionAndRedirect(c *gin.Context, cust *models.Customer) {
	token, err := auth.Sign(auth.SessionUser{
		ID:    cust.ID,
		Name:  cust.Name,
		Email: cust.Email,
		Role:  "customer",
	}, h.secret)
	if err != nil {
		renderAuth(c, http.StatusInternalServerError, "login", gin.H{"Error": "Could not create session. Try again.", "Tab": "login"})
		return
	}
	c.SetCookie(auth.CustomerCookieName, token, 60*60*8, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/me")
}
