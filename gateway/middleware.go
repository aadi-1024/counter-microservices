package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type jsonResponse struct {
	Message string
}

type claims struct {
	Uid int
	Expiry time.Time
	jwt.RegisteredClaims
}

func JWTMiddleware(key []byte) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			//Authorization: Bearer jwt
			token := strings.Split(c.Request().Header.Get("Authorization"), " ")
			if len(token) < 2 {
				return c.JSON(http.StatusUnauthorized, jsonResponse{
					Message: "no JWT token",
				})
			}

			tokenString := token[1]

			clms := &claims{}
			parsedToken, err := jwt.ParseWithClaims(tokenString, clms, func(t *jwt.Token) (interface{}, error) {
				return key, nil
			})

			if !parsedToken.Valid || err != nil {
				return c.JSON(http.StatusUnauthorized, jsonResponse{
					Message: err.Error(),
				})
			}
			
			if clms.Expiry.Before(time.Now()) {
				return c.JSON(http.StatusUnauthorized, jsonResponse{
					Message: "JWT expired",
				})
			}

			c.Set("userid", clms.Uid)
			return next(c)
		}
	}
}