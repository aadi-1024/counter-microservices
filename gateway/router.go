package main

import (
	"counterproto"
	"gateway/handlers"
	"loggerclient"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo) {
	e.Use(loggerclient.LoggerMiddleware(app.Logger))

	e.POST("/login", handlers.LoginHandler(app.Auth))
	e.POST("/register", handlers.RegisterHandler(app.Auth, app.Counter))

	grp := e.Group("/user", JWTMiddleware([]byte("HUGE_SECRET")))
	grp.GET("/get", handlers.GetValue(app.Counter))
	grp.POST("/set", handlers.UpdateValue(app.Counter, counterproto.Action_SetValue))
	grp.POST("/increment", handlers.UpdateValue(app.Counter, counterproto.Action_Increment))
	grp.POST("/decrement", handlers.UpdateValue(app.Counter, counterproto.Action_Decrement))
}
