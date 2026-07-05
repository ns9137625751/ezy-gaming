package customer

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var funcMap = template.FuncMap{
	"add":      func(a, b int) int { return a + b },
	"sub":      func(a, b int) int { return a - b },
	"safeHTML": func(s string) template.HTML { return template.HTML(s) },
	"until": func(n int) []int {
		s := make([]int, n)
		for i := range s {
			s[i] = i
		}
		return s
	},
	"first": func(n int, s string) string {
		r := []rune(s)
		if n > len(r) {
			n = len(r)
		}
		return string(r[:n])
	},
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
}

func injectIsLoggedIn(c *gin.Context, data any) {
	if m, ok := data.(gin.H); ok {
		loggedIn := false
		if v, exists := c.Get("customer_logged_in"); exists {
			if b, ok := v.(bool); ok {
				loggedIn = b
			}
		}
		m["IsLoggedIn"] = loggedIn
	}
}

func renderAuth(c *gin.Context, status int, page string, data any) {
	injectIsLoggedIn(c, data)
	partials, _ := filepath.Glob("templates/partials/*.html")
	files := append([]string{
		"templates/layouts/base.html",
		"templates/customer/" + page + ".html",
	}, partials...)

	tmpl, err := template.New("").Funcs(funcMap).ParseFiles(files...)
	if err != nil {
		log.Printf("customer auth template parse error (%s): %v", page, err)
		c.String(http.StatusInternalServerError, "template error")
		return
	}
	c.Status(status)
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(c.Writer, "base", data); err != nil {
		log.Printf("customer auth template execute error (%s): %v", page, err)
	}
}

func renderPage(c *gin.Context, status int, page string, data any) {
	injectIsLoggedIn(c, data)
	partials, _ := filepath.Glob("templates/partials/*.html")
	files := append([]string{
		"templates/layouts/base.html",
		"templates/customer/" + page + ".html",
	}, partials...)

	tmpl, err := template.New("").Funcs(funcMap).ParseFiles(files...)
	if err != nil {
		log.Printf("customer template parse error (%s): %v", page, err)
		c.String(http.StatusInternalServerError, "template error")
		return
	}
	c.Status(status)
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(c.Writer, "base", data); err != nil {
		log.Printf("customer template execute error (%s): %v", page, err)
	}
}
