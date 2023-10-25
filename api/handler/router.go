package handler

import (
	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, petHandler PetHandler) {
	e.GET("/pet/:id", petHandler.TopView())
	e.POST("pet/:id/schedule", petHandler.PostSchedule())
}
