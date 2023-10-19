package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	dbconnect "github.com/myantyuWorld/animal_healthcate/database"
)

type Pets struct {
	Name      string
	BirthDay  time.Time
	Sex       int8 // 0:オス, 1: メス
	CreatedAt time.Time
	UpdatedAt time.Time
}

func seeds(db *gorm.DB) error {
	// TODO: 日本語を投入しようとすると以下警告
	// Error 1366 (HY000): Incorrect string value: '\xE3\x83\xA1\xE3\x82\xB9' for column 'sex' at row 1root@3982857c066f:/go/src/app/seed#
	pet := Pets{Name: "natsu", BirthDay: time.Now(), Sex: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	if err := db.Create(&pet).Error; err != nil {
		fmt.Printf("%+v", err)
	}
	pet1 := Pets{Name: "chai", BirthDay: time.Now(), Sex: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	if err := db.Create(&pet1).Error; err != nil {
		fmt.Printf("%+v", err)
	}
	pet2 := Pets{Name: "mugi", BirthDay: time.Now(), Sex: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	if err := db.Create(&pet2).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	return nil
}

func main() {
	// TODO: database.connect.goのConnect()でできるので上記メソッドがいらない
	// db := dbconnect.Connect()
	db := dbconnect.Connect()
	if err := seeds(db); err != nil {
		fmt.Printf("%+v", err)
		return
	}
	defer db.Close()
}
