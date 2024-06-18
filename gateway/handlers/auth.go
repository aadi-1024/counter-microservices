package handlers

import (
	"authproto"
	"context"
	"counterproto"
	"gateway/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginHandler(client authproto.AuthRPCClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		loginReq := authproto.LoginRequest{}
		loginReq.Email = c.FormValue("email")
		loginReq.Password = c.FormValue("password")

		resp, err := client.Login(c.Request().Context(), &loginReq)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		jsonResp := &models.JsonResponse{}
		jsonResp.Message = resp.Message

		jsonResp.Token = resp.Token
		return c.JSON(http.StatusOK, jsonResp)
	}
}

func RegisterHandler(client authproto.AuthRPCClient, cclient counterproto.CounterRPCClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		registerReq := authproto.RegisterRequest{}
		registerReq.Email = c.FormValue("email")
		registerReq.Username = c.FormValue("username")
		registerReq.Password = c.FormValue("password")

		resp, err := client.Register(c.Request().Context(), &registerReq)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		_, err = cclient.CreateNew(context.Background(), &counterproto.Request{
			UserId: resp.Userid,
			Value:  0,
		})

		jsonResp := &models.JsonResponse{}
		if err != nil {
			jsonResp.Message = err.Error()
			return c.JSON(http.StatusBadRequest, jsonResp)
		}
		jsonResp.Message = "successful"
		jsonResp.Token = resp.Token

		return c.JSON(http.StatusOK, jsonResp)
	}
}
