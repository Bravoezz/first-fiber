package db

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/Bravoezz/first-fiber/config"
	"github.com/Bravoezz/first-fiber/modules/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var once sync.Once
var DB *gorm.DB

func InitDbConnection() {
	var (
		dbname        = config.GetEnv("DB_NAME")
		dbuser        = config.GetEnv("DB_USER")
		dbpassword    = config.GetEnv("DB_PASSWORD")
		dbhost        = config.GetEnv("DB_HOST")
		dbport        = config.GetEnv("DB_PORT")
		uriConnection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbport, dbname)
	)
	var err error

	once.Do(func() {
		DB, err = gorm.Open(mysql.Open(uriConnection), &gorm.Config{ Logger: NewLogger()})

		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("DB Connected")
		}
	})
}

func NewLogger() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)
	return newLogger
}

func Migrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		fmt.Println("Error en la migracion de modelos")
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("Migracion correcta")
	}
}
