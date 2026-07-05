package renderer

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var partialFiles []string

// funcMap provides custom template helpers not built into html/template.
var funcMap = template.FuncMap{
	// until returns a slice of ints [0, n) — used for star ratings, repeat loops.
	"until": func(n int) []int {
		s := make([]int, n)
		for i := range s {
			s[i] = i
		}
		return s
	},
	// first returns the first n runes of a string.
	"first": func(n int, s string) string {
		r := []rune(s)
		if n > len(r) {
			n = len(r)
		}
		return string(r[:n])
	},
	// percent returns floor(a*100/b), clamped to [0,100]. Returns 0 if b==0.
	"percent": func(a, b int) int {
		if b == 0 {
			return 0
		}
		v := a * 100 / b
		if v > 100 {
			return 100
		}
		return v
	},
	// safeHTML passes a string through as template.HTML (used for JSON-LD blobs).
	"safeHTML": func(s string) template.HTML {
		return template.HTML(s)
	},
	// add performs integer addition (used in templates for indexing).
	"add": func(a, b int) int { return a + b },
	// sub performs integer subtraction.
	"sub": func(a, b int) int { return a - b },
}

func init() {
	partials, err := filepath.Glob("templates/partials/*.html")
	if err != nil || len(partials) == 0 {
		log.Println("Warning: no partial templates found")
	}
	partialFiles = partials
}

// Render parses base layout + page + partials on each request.
// Each call builds an independent template set so "content" blocks
// don't conflict across pages.
func Render(c *gin.Context, status int, page string, data any) {
	// Inject IsLoggedIn so the shared navbar can hide the Login button
	if m, ok := data.(gin.H); ok {
		loggedIn := false
		if v, exists := c.Get("customer_logged_in"); exists {
			if b, ok := v.(bool); ok {
				loggedIn = b
			}
		}
		m["IsLoggedIn"] = loggedIn
	}

	files := []string{
		"templates/layouts/base.html",
		"templates/pages/" + page + ".html",
	}
	files = append(files, partialFiles...)

	tmpl, err := template.New("").Funcs(funcMap).ParseFiles(files...)
	if err != nil {
		log.Printf("Template parse error (%s): %v", page, err)
		c.String(http.StatusInternalServerError, "template error: %v", err)
		return
	}

	c.Status(status)
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(c.Writer, "base", data); err != nil {
		log.Printf("Template execute error (%s): %v", page, err)
	}
}

// isDev returns true when not in production.
func isDev() bool {
	return os.Getenv("APP_ENV") != "production"
}
