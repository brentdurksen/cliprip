package main

type createData struct {
	Key string
}

func (app *application) routes() {
	app.router.Static("/static", "./ui/static")
	app.router.Get("/", app.getHome)
	// Clipboard routes
	clipboard := app.router.Group("/clipboard")
	{
		clipboard.Get("/:key", app.getClipboard)
		clipboard.Post("/save", app.saveClipboard)
		clipboard.Delete("/:key", app.deleteClipboard)
	}
	// Components
	components := app.router.Group("/components")
	{
		components.Get("/clipboard", app.getClipboardComponent)
	}
}
