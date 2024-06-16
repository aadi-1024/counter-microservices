package main

import (
	"gateway/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo) {
	e.POST("/login", handlers.LoginHandler(app.Auth))
	// e.POST("/register", nil)
}
