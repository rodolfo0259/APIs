package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type LocationsHandler struct{}

func (h *LocationsHandler) All(c echo.Context) error {
	return c.JSON(http.StatusOK, []map[string]string{
		{"getting": "started"},
	})

}
