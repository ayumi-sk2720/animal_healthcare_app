package dbconnect

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// TODO: seedで使っているが、sqlhandlerを使うように変更したい
func Connect() *gorm.DB {
	dsn := "root:hogehoge@tcp(db:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	return db
}
