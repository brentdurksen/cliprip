package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
)

func (app *application) getHome(c *gin.Context) {
	b := new(bytes.Buffer)
	app.templates.ExecuteTemplate(b, "search.tmpl", nil)
	app.templates.ExecuteTemplate(c.Writer, "base.tmpl", template.HTML(b.String()))
}

func (app *application) getClipboardComponent(c *gin.Context) {
	key := c.Request.FormValue("key")
	// validate key is a SHA-256 hash
	validators.validateHash(key)

	c.Header("HX-Push-Url", fmt.Sprintf("/clipboard/%s", key))
	if _, ok := app.secrets[key]; ok {
		app.templates.ExecuteTemplate(c.Writer, "read.tmpl", app.secrets[key])
	} else {
		app.templates.ExecuteTemplate(c.Writer, "create.tmpl", createData{Key: c.Request.FormValue("key")})
	}
}

func (app *application) getReadComponent(c *gin.Context) {
	key := c.Param("key")
	if _, ok := app.secrets[key]; !ok {
		app.notFound(c)
		return
	}
	app.templates.ExecuteTemplate(c.Writer, "read.tmpl", app.secrets[key])
}

func (app *application) getClipboard(c *gin.Context) {
	key := c.Param("key")
	if _, ok := app.secrets[key]; !ok {
		app.notFound(c)
		return
	}
	b := new(bytes.Buffer)
	app.templates.ExecuteTemplate(b, "read.tmpl", app.secrets[key])
	app.templates.ExecuteTemplate(c.Writer, "base.tmpl", template.HTML(b.String()))
}

func (app *application) saveClipboard(c *gin.Context) {
	key := c.Request.FormValue("key")
	value := c.Request.FormValue("content")
	app.secrets[key] = value
	app.templates.ExecuteTemplate(c.Writer, "read.tmpl", value)
}

func (app *application) notFound(c *gin.Context) {
	b := new(bytes.Buffer)
	app.templates.ExecuteTemplate(b, "404.tmpl", nil)
	app.templates.ExecuteTemplate(c.Writer, "base.tmpl", template.HTML(b.String()))
	c.AbortWithStatus(404)
}
