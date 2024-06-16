package handlers

import (
	"authproto"
	"net/http"

	"github.com/labstack/echo/v4"
)

type jsonResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func LoginHandler(client authproto.AuthRPCClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		loginReq := authproto.LoginRequest{}
		loginReq.Email = c.FormValue("email")
		loginReq.Password = c.FormValue("password")

		resp, err := client.Login(c.Request().Context(), &loginReq)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		jsonResp := &jsonResponse{}
		jsonResp.Message = resp.Message

		if !resp.Success {
			return c.JSON(http.StatusBadRequest, jsonResp)
		}

		jsonResp.Token = resp.Token
		return c.JSON(http.StatusOK, jsonResp)
	}
}
