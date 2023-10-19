package main

import (
	"net/http"
	// controller "api/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/myantyuWorld/animal_healthcate/controllers"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORS())

	// ルーティング
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, golang echo!")
	})
	e.GET("pet/:id", controllers.FetchPet())
	e.GET("pet/:id/schedule", controllers.GetShedules())
	return e
}
