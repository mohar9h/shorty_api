package web

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"path/filepath"
)

func SubmitPage(c *gin.Context) {
	tmplPath := filepath.Join("internal", "web", "templates", "submit.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Template parse error: %v", err)
		return
	}
	_ = tmpl.Execute(c.Writer, nil)
}

func ResultPage(c *gin.Context) {
	code := c.Param("code")
	data := map[string]interface{}{
		"Code":    code,
		"FullURL": c.Request.Host + "/" + code,
	}

	tmplPath := filepath.Join("internal", "web", "templates", "result.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Template parse error: %v", err)
		return
	}
	_ = tmpl.Execute(c.Writer, data)
}
