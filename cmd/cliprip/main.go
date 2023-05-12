package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
)

type application struct {
	router  *fiber.App
	secrets map[string]string // replace with SQLite
}

func main() {
	engine := html.New("./ui/html", ".tmpl")
	f := fiber.New(fiber.Config{
		Views:        engine,
		ErrorHandler: errorHandler,
	})
	f.Use(recover.New())
	f.Use(compress.New())
	f.Use(favicon.New(favicon.Config{File: "./ui/static/favicon.ico"}))

	secrets := make(map[string]string)

	app := &application{f, secrets}
	app.routes()
	app.router.Listen(":8080")
}
