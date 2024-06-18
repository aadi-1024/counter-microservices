package handlers

import (
	"counterproto"
	"gateway/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UpdateValue(client counterproto.CounterRPCClient, act counterproto.Action) echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("userid").(int)
		value, err := strconv.Atoi(c.FormValue("value"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, &models.JsonResponse{
				Message: err.Error(),
			})
		}

		req := &counterproto.Request{
			UserId: int32(uid),
			Value: int32(value),
			Task: act,
		}

		resp, err := client.Update(c.Request().Context(), req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &models.JsonResponse{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: resp.Message,
			Value: int(resp.Value),
		})
	}
}

func GetValue(client counterproto.CounterRPCClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("userid").(int)
		
		req := &counterproto.Request{
			UserId: int32(uid),
		}

		resp, err := client.GetValue(c.Request().Context(), req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &models.JsonResponse{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: resp.Message,
			Value: int(resp.Value),
		})
	}
}