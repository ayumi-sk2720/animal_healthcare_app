package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	dbconnect "github.com/myantyuWorld/animal_healthcate/database"
)

func FetchPet() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameter取得
		id := c.Param("id")
		fmt.Printf("pet id:: %v\n", id)

		// db接続を取得し、deferで閉じる
		db := dbconnect.Connect()
		defer db.Close()

		// ORM

		// db操作

		return c.JSON(http.StatusOK, "get pet")
	}
}

func PostShedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameter取得

		// db接続を取得し、deferで閉じる

		// ORM

		// db操作
		return c.JSON(http.StatusOK, "post shedule!")
	}
}
