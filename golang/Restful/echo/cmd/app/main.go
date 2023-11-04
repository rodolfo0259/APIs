package main

import (
	"echo_restful/internal/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()
	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	api := router.Group("")
	{
		routes.RegisterLocations(api)
	}

	router.Logger.Fatal(router.Start(":1323"))

}
