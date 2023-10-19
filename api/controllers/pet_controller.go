package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	dbconnect "github.com/myantyuWorld/animal_healthcate/database"
	"github.com/myantyuWorld/animal_healthcate/model"
)

// TODO: Golangで、デザインパターン「Template Method」を学ぶ | https://qiita.com/ttsubo/items/a36c6fc9acac3513b4af
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
		// TODO: 将来的には、テーブル結合した結果を返却する
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

// TODO: クエリパラメータを指定して、期間を絞り込んだり、当日のデータを返すようなものが望ましい
func GetShedules() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameter取得
		id := c.Param("id")

		// db接続を取得し、deferで閉じる
		db := dbconnect.Connect()
		defer db.Close()

		// ORM
		schedules := []model.Schedule{}
		result := db.Find(&schedules, "pet_id = ?", id)

		if result.RecordNotFound() {
			fmt.Printf("レコードが見つかりません")
			// TODO: エラーを共通的なJSONで返したい
			return c.JSON(http.StatusFound, "該当のペットに対応するスケジュールが見つかりませんでした。")
		}

		// db操作
		return c.JSON(http.StatusOK, &schedules)
	}
}
