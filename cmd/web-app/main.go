package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

type application struct {
	router    *gin.Engine
	templates *template.Template
	secrets   map[string]string // replace with database
}

func main() {
	router := gin.Default()
	templates := template.Must(template.ParseGlob("ui/html/*.tmpl"))
	secrets := make(map[string]string)

	app := &application{router, templates, secrets}
	app.routes()
	app.router.Run()
}
