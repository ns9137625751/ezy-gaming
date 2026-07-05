package admin

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var funcMap = template.FuncMap{
	"add": func(a, b int) int { return a + b },
	"first": func(n int, s string) string {
		r := []rune(s)
		if n > len(r) {
			n = len(r)
		}
		return string(r[:n])
	},
}

func renderAdmin(c *gin.Context, status int, page string, data any) {
	tmpl, err := template.New("").Funcs(funcMap).ParseFiles(
		"templates/admin/layout.html",
		"templates/admin/"+page+".html",
	)
	if err != nil {
		log.Printf("admin template parse error (%s): %v", page, err)
		c.String(http.StatusInternalServerError, "template error")
		return
	}
	c.Status(status)
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(c.Writer, "admin_layout", data); err != nil {
		log.Printf("admin template execute error (%s): %v", page, err)
	}
}

func renderLogin(c *gin.Context, status int, data any) {
	tmpl, err := template.New("").Funcs(funcMap).ParseFiles("templates/admin/login.html")
	if err != nil {
		log.Printf("admin login template parse error: %v", err)
		c.String(http.StatusInternalServerError, "template error")
		return
	}
	c.Status(status)
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(c.Writer, "admin_login", data); err != nil {
		log.Printf("admin login template execute error: %v", err)
	}
}
