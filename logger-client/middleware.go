package loggerclient

import "github.com/labstack/echo/v4"

//Only for Echo
func LoggerMiddleware(l *Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			msg := "Received " + c.Request().Method + "Request at " + c.Request().URL.RawPath
			l.Log(msg, Info)
			return next(c)
		}
	}
}