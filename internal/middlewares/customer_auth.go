package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/auth"
)

func CustomerAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tok, err := c.Cookie(auth.CustomerCookieName)
		if err != nil || tok == "" {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}
		user, err := auth.Verify(tok, secret)
		if err != nil {
			c.SetCookie(auth.CustomerCookieName, "", -1, "/", "", false, true)
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}
		c.Set("customer", user)
		c.Next()
	}
}

// CustomerSessionInfo is a non-blocking middleware that detects whether a
// customer is logged in and stores the result in the gin context as
// "customer_logged_in". Used by public pages to hide/show the Login button.
func CustomerSessionInfo(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tok, err := c.Cookie(auth.CustomerCookieName)
		if err == nil && tok != "" {
			if _, err := auth.Verify(tok, secret); err == nil {
				c.Set("customer_logged_in", true)
				c.Next()
				return
			}
		}
		c.Set("customer_logged_in", false)
		c.Next()
	}
}
