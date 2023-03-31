package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed static/*
var staticFS embed.FS

func main() {

	r := gin.Default()

	tmpl := template.Must(template.New("").ParseFS(staticFS, "static/*/*"))
	r.SetHTMLTemplate(tmpl)

	r.GET("/", func(c *gin.Context) { c.Redirect(http.StatusPermanentRedirect, "/index") })

	r.GET("/index", func(c *gin.Context) {
		r.LoadHTMLGlob("static/*/*")
		c.HTML(http.StatusOK, "page.tmpl.html", gin.H{
			"Template": "index",
			"Title":    "Main website",
		})
	})

	r.GET("/A", func(c *gin.Context) {
		r.LoadHTMLGlob("static/*/*")
		c.HTML(http.StatusOK, "page.tmpl.html", gin.H{
			"Template": "a",
			"Title":    "A website",
		})
	})

	r.GET("/B", func(c *gin.Context) {
		r.LoadHTMLGlob("static/*/*")
		c.HTML(http.StatusOK, "page.tmpl.html", gin.H{
			"Template": "b",
			"Title":    "B website",
		})
	})

	r.Run(":9999")
}
