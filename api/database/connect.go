package dbconnect

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// TODO: seedで使っているが、sqlhandlerを使うように変更したい
func Connect() *gorm.DB {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST") // ココ!!
	dbname := os.Getenv("MYSQL_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	return db
}
