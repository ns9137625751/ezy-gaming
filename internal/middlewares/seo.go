package middlewares

import "github.com/gin-gonic/gin"

// SEOHeaders injects HTTP headers that improve security scores,
// tell crawlers to trust the canonical origin, and enable
// browser-level performance optimisations.
func SEOHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Security headers (improve Google trust score)
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "SAMEORIGIN")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")

		// Tell crawlers the canonical domain
		c.Header("X-Robots-Tag", "index, follow")

		// Enable Gzip via Content-Encoding negotiation
		c.Header("Vary", "Accept-Encoding")

		// Static asset caching (handled separately for /static/ routes)
		// For HTML pages: no-store so Google always gets fresh content
		path := c.Request.URL.Path
		switch {
		case isStaticPath(path):
			// 1 year cache for fingerprinted assets
			c.Header("Cache-Control", "public, max-age=31536000, immutable")
		case path == "/sitemap.xml":
			c.Header("Cache-Control", "public, max-age=3600") // re-crawl hourly
		default:
			// HTML pages: private cache, revalidate
			c.Header("Cache-Control", "no-cache, must-revalidate")
		}

		c.Next()
	}
}

func isStaticPath(path string) bool {
	return len(path) > 8 && path[:8] == "/static/"
}
