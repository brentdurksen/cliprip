package main

type createData struct {
	Key string
}

func (app *application) routes() {
	app.router.Static("/static", "ui/static")
	app.router.NoRoute(app.notFound)
	app.router.GET("/", app.getHome)
	app.router.StaticFile("/favicon.ico", "ui/static/favicon.ico")
	// Clipboard routes
	clipboard := app.router.Group("/clipboard")
	{
		clipboard.GET("/:key", app.getClipboard)
		clipboard.POST("/save", app.saveClipboard)
		clipboard.POST("/delete", app.deleteClipboard)
	}
	// Components
	components := app.router.Group("/components")
	{
		components.GET("/clipboard", app.getClipboardComponent)
	}
}
