package main

import (
	"fmt"
	"net/http"

	// controller "api/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/myantyuWorld/animal_healthcate/controllers"
)

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("Request Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
	fmt.Println("==================================================================")
}

func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.BodyDump(bodyDumpHandler))

	// ルーティング
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, golang echo!")
	})
	e.GET("pet/:id", controllers.FetchPet())
	e.GET("pet/:id/schedules", controllers.GetShedules())
	e.POST("pet/:id/schedules", controllers.PostSchedule())

	return e
}
