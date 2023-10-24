package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	dbconnect "github.com/myantyuWorld/animal_healthcate/database"
)

type Pets struct {
	Name         string
	BirthDay     time.Time
	Sex          int8 // 0:オス, 1: メス
	NowWeight    int
	TargetWeight int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func seeds(db *gorm.DB) error {
	pet := Pets{
		Name:         "なっちゃん",
		BirthDay:     time.Now(),
		Sex:          1,
		NowWeight:    4000,
		TargetWeight: 3500,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := db.Create(&pet).Error; err != nil {
		fmt.Printf("%+v", err)
	}
	pet1 := Pets{
		Name:         "チャイ",
		BirthDay:     time.Now(),
		Sex:          1,
		NowWeight:    4000,
		TargetWeight: 9999,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := db.Create(&pet1).Error; err != nil {
		fmt.Printf("%+v", err)
	}
	pet2 := Pets{
		Name:         "むぎ",
		BirthDay:     time.Now(),
		Sex:          1,
		NowWeight:    4000,
		TargetWeight: 9999,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := db.Create(&pet2).Error; err != nil {
		fmt.Printf("%+v", err)
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
