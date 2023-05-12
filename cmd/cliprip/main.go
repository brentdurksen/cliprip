package main

import (
	"html/template"

	"github.com/gin-contrib/gzip"
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
)

type application struct {
	router    *gin.Engine
	templates *template.Template
	secrets   map[string]string // replace with SQLite
}

func main() {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(limits.RequestSizeLimiter(256000))
	templates := template.Must(template.ParseGlob("ui/html/*.tmpl"))
	secrets := make(map[string]string)

	app := &application{r, templates, secrets}
	app.routes()
	app.router.Run()
}
