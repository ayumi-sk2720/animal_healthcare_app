package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	dbconnect "github.com/myantyuWorld/animal_healthcate/database"
	"github.com/myantyuWorld/animal_healthcate/domain/model"
)

// TODO: [同じ処理が多い、テンプレートパターンで処理を抽象化したい] Golangで、デザインパターン「Template Method」を学ぶ | https://qiita.com/ttsubo/items/a36c6fc9acac3513b4af
// TODO: [golang Clean Architecture] | https://qiita.com/mIchino/items/b885de3396e3f77d8b37
func FetchPet() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameter取得
		id := c.Param("id")
		fmt.Printf("pet id:: %v\n", id)

		// db接続を取得し、deferで閉じる
		db := dbconnect.Connect()
		defer db.Close()

		// ORM
		topInfo := model.TopInfo{}
		pet := model.Pet{}

		// db操作 | https://gorm.io/ja_JP/docs/query.html

		result := db.First(&pet, id)
		fmt.Printf("pet:: %v\n", &pet)
		topInfo.Pet = pet
		// TODO: 将来的には、テーブル結合した結果を返却する --- start ---
		topInfo.DosageSchedules.Today = model.Task{
			Title: "hogehoge",
			Date:  time.Now(),
		}
		topInfo.DosageSchedules.Next = model.Task{
			Title: "fugafuga",
			Date:  time.Now(),
		}
		topInfo.PhysicalCondition.Food = 3
		topInfo.PhysicalCondition.Sweat = 2
		topInfo.PhysicalCondition.Toilet = 1
		topInfo.Memo = model.Task{
			Title: "今月に入って飲む水の量が\n増えた気がする\n次に通院した時に先生に相談する",
			Date:  time.Now(),
		}
		topInfo.Schedules = []model.Schedule{
			{
				PetId:    id,
				Title:    "トリミング",
				Date:     time.Now(),
				Location: "ペテモ立川店",
			},
			{
				PetId:    id,
				Title:    "通院",
				Date:     time.Now(),
				Location: "ホゲホゲ病院",
			},
		}
		// TODO: 将来的には、テーブル結合した結果を返却する --- end ---

		if result.RecordNotFound() {
			fmt.Printf("レコードが見つかりません")
			// TODO: エラーを共通的なJSONで返したい
			return c.JSON(http.StatusFound, "該当のペットが見つかりませんでした")
		}

		return c.JSON(http.StatusOK, topInfo)
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

// [example]
//
//	curl "http://0.0.0.0:8080/pet/1/schedules" \
//	  -X POST \
//	  -d '{"title": "サンプル0",    "date": "2023-10-19T18:58:18+09:00",    "location": "サンプルペットショップ0"}' \
//	  -H "content-type: application/json"
func PostSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Binding | https://echo.labstack.com/docs/binding
		schedule := new(model.Schedule)
		if error := c.Bind(&schedule); error != nil {
			fmt.Println()
			return c.JSON(http.StatusBadRequest, "bad request")
		}
		schedule.PetId = c.Param("id")
		schedule.CreatedAt = time.Now()
		schedule.UpdatedAt = time.Now()
		fmt.Println(schedule)

		// db接続を取得し、deferで閉じる
		db := dbconnect.Connect()
		defer db.Close()

		if error := db.Create(schedule).Error; error != nil {
			return c.JSON(http.StatusInternalServerError, "異常が発生しました、担当者に連絡してください。")
		}

		return c.JSON(http.StatusOK, &schedule)
	}
}
