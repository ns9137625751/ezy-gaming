package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/auth"
)

const adminUserKey = "admin_user"

// AdminAuth validates the signed session cookie and puts the user into context.
func AdminAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(auth.CookieName)
		if err != nil || token == "" {
			c.Redirect(http.StatusSeeOther, "/admin/login")
			c.Abort()
			return
		}
		user, err := auth.Verify(token, secret)
		if err != nil {
			c.SetCookie(auth.CookieName, "", -1, "/", "", false, true)
			c.Redirect(http.StatusSeeOther, "/admin/login")
			c.Abort()
			return
		}
		c.Set(adminUserKey, user)
		c.Next()
	}
}

// RequireRole aborts with 403 if the session user's role is not in the allowed list.
func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, ok := c.Get(adminUserKey)
		if !ok {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		user := u.(*auth.SessionUser)
		for _, r := range roles {
			if user.Role == r {
				c.Next()
				return
			}
		}
		c.AbortWithStatus(http.StatusForbidden)
	}
}

// GetAdminUser is a helper for controllers to extract the session user from context.
func GetAdminUser(c *gin.Context) *auth.SessionUser {
	u, _ := c.Get(adminUserKey)
	if u == nil {
		return nil
	}
	return u.(*auth.SessionUser)
}
