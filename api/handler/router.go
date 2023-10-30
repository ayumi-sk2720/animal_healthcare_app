package handler

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// middleware
func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("Request Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
	fmt.Println("==================================================================")
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func InitRouting(e *echo.Echo, petHandler PetHandler) {
	e.Use(middleware.Logger())
	e.Use(middleware.CORS()) // 以下を追加
	e.Use(middleware.BodyDump(bodyDumpHandler))

	// validatorを登録
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/pet/:id", petHandler.TopView())
	e.POST("/pet/:id/schedule", petHandler.PostSchedule())
}
