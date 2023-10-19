package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	dbconnect "github.com/myantyuWorld/animal_healthcate/database"
	"github.com/myantyuWorld/animal_healthcate/model"
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
		pet := model.Pet{}

		// db操作 | https://gorm.io/ja_JP/docs/query.html
		result := db.First(&pet, id)
		fmt.Printf("pet:: %v\n", &pet)

		if result.RecordNotFound() {
			fmt.Printf("レコードが見つかりません")
			// TODO: エラーを共通的なJSONで返したい
			return c.JSON(http.StatusFound, "該当のペットが見つかりませんでした")
		}

		return c.JSON(http.StatusOK, pet)
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
