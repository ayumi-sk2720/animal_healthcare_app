package handler

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// middleware
func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("Request Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
	fmt.Println("==================================================================")
}

func InitRouting(e *echo.Echo, petHandler PetHandler) {
	e.Use(middleware.CORS())
	e.Use(middleware.BodyDump(bodyDumpHandler))

	e.GET("/pet/:id", petHandler.TopView())
	e.POST("pet/:id/schedule", petHandler.PostSchedule())
}
