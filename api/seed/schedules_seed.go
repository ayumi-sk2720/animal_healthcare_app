package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	dbconnect "github.com/myantyuWorld/animal_healthcate/database"
)

type Schedules struct {
	PetId     int
	Title     string
	Date      time.Time
	Location  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func seeds(db *gorm.DB) error {
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			schedule := Schedules{
				PetId:     i + 1,
				Title:     "サンプル" + strconv.Itoa(i),
				Date:      time.Now(),
				Location:  "サンプルペットショップ" + strconv.Itoa(j),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := db.Create(&schedule).Error; err != nil {
				fmt.Printf("%+v", err)
			}
		}
	}

	return nil
}

func main() {
	db := dbconnect.Connect()
	if err := seeds(db); err != nil {
		fmt.Printf("%+v", err)
		return
	}
	defer db.Close()
}
