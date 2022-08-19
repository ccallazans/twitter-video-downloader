package main

import (
	"github.com/labstack/echo/v4"
)

func (app *Application) NewRouter() {
	router := echo.New()

	router.GET("/:url", app.GetUrl)

	app.server = router
}
