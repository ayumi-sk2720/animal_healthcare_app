package handler

import (
	"fmt"
	"net/http"

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
	e.Use(middleware.Logger())
	e.Use(middleware.CORS()) // 以下を追加
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", "http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
	e.Use(middleware.BodyDump(bodyDumpHandler))

	e.GET("/pet/:id", petHandler.TopView())
	e.POST("pet/:id/schedule", petHandler.PostSchedule())
}
