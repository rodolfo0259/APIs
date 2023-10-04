package routes

import (
	"echo_restful/internal/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterLocations(router *echo.Group) {
	handler := handlers.LocationsHandler{}

	v1 := router.Group("/locations")
	{
		v1.GET("", handler.All)
	}
}
