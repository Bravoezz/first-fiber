package db

import (
	"fmt"
	"log"
	"sync"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once
var DB *gorm.DB

func InitDbConnection() {
	var (
		dbname        = "fiberapi"
		dbuser        = "root"
		dbpassword    = "toor"
		dbhost        = "localhost"
		dbport        = "3306"
		uriConnection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbport, dbname)
	)
	var err error

	once.Do(func() {
		DB, err = gorm.Open(mysql.Open(uriConnection), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}else {
			log.Println("DB Connected")
		}
	})
}
