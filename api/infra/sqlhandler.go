package infra

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SqlHandler struct {
	Conn *gorm.DB
}

var count = 30

func connect(path string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", path)
	if err != nil {
		log.Println("Not ready. Retry connecting...")
		time.Sleep(time.Second)
		count++
		log.Println(count)
		if count > 30 {
			panic(err)
		}
		return connect(path)
	}
	log.Println("Successfully")
	return db, nil
}

func NewSqlHandler() *SqlHandler {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST") // ココ!!
	dbname := os.Getenv("MYSQL_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)

	// conn, err := gorm.Open("mysql", dsn)
	conn, err := connect(dsn)
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
