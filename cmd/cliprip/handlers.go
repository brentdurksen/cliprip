package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func (app *application) getHome(c *fiber.Ctx) error {
	return c.Render("search", nil, "base")
}

func (app *application) getClipboardComponent(c *fiber.Ctx) error {
	key := c.FormValue("key")
	if ok := validateHash(key); !ok {
		return c.SendStatus(http.StatusBadRequest)
	}

	c.Set("HX-Push-Url", fmt.Sprintf("/clipboard/%s", key))
	if _, ok := app.secrets[key]; ok {
		return c.Render("read", app.secrets[key])
	} else {
		return c.Render("create", &fiber.Map{"Key": key})
	}
}

func (app *application) getReadComponent(c *fiber.Ctx) error {
	key := c.Params("key")
	if ok := validateHash(key); !ok {
		return app.badrequest(c)
	}
	if _, ok := app.secrets[key]; !ok {
		return app.notFound(c)
	}
	return c.Render("read", app.secrets[key])
}

func (app *application) getClipboard(c *fiber.Ctx) error {
	key := c.Params("key")
	if ok := validateHash(key); !ok {
		return app.badrequest(c)
	}
	if _, ok := app.secrets[key]; !ok {
		return app.notFound(c)
	}
	return c.Render("read", app.secrets[key], "base")
}

func (app *application) saveClipboard(c *fiber.Ctx) error {
	// Fiber reuses contexts, so we need to copy the key
	key := utils.CopyString(c.FormValue("key"))
	if ok := validateHash(key); !ok {
		return c.SendStatus(http.StatusBadRequest)
	}
	// Fiber reuses contexts, so we need to copy the content
	content := utils.CopyString(c.FormValue("content"))
	if len(content) == 0 {
		return c.SendStatus(http.StatusBadRequest)
	}
	app.secrets[key] = content
	return c.Render("read", content)
}

func (app *application) deleteClipboard(c *fiber.Ctx) error {
	key := c.Params("key")
	if ok := validateHash(key); !ok {
		return app.badrequest(c)
	}
	delete(app.secrets, key)
	return app.getClipboardComponent(c)
}

func (app *application) notFound(c *fiber.Ctx) error {
	return c.Status(http.StatusNotFound).Render("error", &fiber.Map{"ErrorCode": "404", "ErrorMessage": "Not Found"}, "base")
}
func (app *application) badrequest(c *fiber.Ctx) error {
	return c.Status(http.StatusBadRequest).Render("error", &fiber.Map{"ErrorCode": "400", "ErrorMessage": "Bad Request"}, "base")
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	return ctx.Status(code).Render("error", &fiber.Map{"ErrorCode": fmt.Sprintf("%v", code)}, "base")
}
