package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	dbconnect "github.com/myantyuWorld/animal_healthcate/database"
	"github.com/myantyuWorld/animal_healthcate/model"
)

// TODO: [同じ処理が多い、テンプレートパターンで処理を抽象化したい] Golangで、デザインパターン「Template Method」を学ぶ | https://qiita.com/ttsubo/items/a36c6fc9acac3513b4af
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
		// 実際にはこのJSONを返す必要がある
		// {
		// 	"baseInfo": {
		// 		"name": "チャイ",
		// 		"age": "3歳8ヶ月",
		// 		"sex": "メス",
		// 		"birthday": "2020年1月27日"
		// 	},
		// 	"now_wight": {
		// 		"weight": "4.8kg",
		// 		"date": "2023年9月13日 10:00"
		// 	},
		// 	"target_wight": {
		// 		"weight": "5.2kg"
		// 	},
		// 	"dosage_schedule": {
		// 		"today": "なし",
		// 		"next": {
		// 			"name": "ネクスがーど",
		// 			"date": "2023年9月15日 10:00"
		// 		}
		// 	},
		// 	"physical_condition": {
		// 		"food": 3,
		// 		"sweat": 2,
		// 		"toilet": 1
		// 	},
		// 	"memo": "今月に入って飲む水の量が\\n増えた気がする\\n次に通院した時に先生に相談する",
		// 	"schedules": [
		// 		{
		// 			"date": "2023年9月20日 10:00",
		// 			"name": "トリミング"
		// 		},
		// 		{
		// 			"date": "2023年9月30日 14:00",
		// 			"name": "通院"
		// 		}
		// 	]
		// }

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

// [example]
//
//	{
//			"title": "サンプル0",
//			"date": "2023-10-19T18:58:18+09:00",
//			"location": "サンプルペットショップ0"
//	}
func PostSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Binding | https://echo.labstack.com/docs/binding
		schedule := new(model.PostSchedule)
		if error := c.Bind(&schedule); error != nil {
			fmt.Println()
			return c.JSON(http.StatusBadRequest, "bad request")
		}
		fmt.Println(schedule)

		return c.JSON(http.StatusOK, &schedule)
	}
}
