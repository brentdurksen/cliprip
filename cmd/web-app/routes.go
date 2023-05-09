package main

type createData struct {
	Key string
}

func (app *application) routes() {
	r := app.router
	r.Static("/static", "ui/static")
	r.NoRoute(app.notFound)
	r.GET("/", app.getHome)
	// Components
	components := r.Group("/components")
	{
		components.GET("/clipboard", app.getClipboardComponent)
	}
	// Pages
	clipboard := r.Group("/clipboard")
	{
		clipboard.GET("/:key", app.getClipboard)
		clipboard.POST("/save", app.saveClipboard)
	}
}
