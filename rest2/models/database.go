package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

// Дескриптор БД
var db *gorm.DB

func init() {
	// Загрузка переменных окружения
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Данные для подключения к БД
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	database := os.Getenv("db_name")
	host := os.Getenv("db_host")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		host, username, password, database,
	)

	// Подключение к БД
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Миграции
	db.Debug().AutoMigrate(&Account{}, &Contact{})
}

// Получение дескриптора БД
func DB() *gorm.DB {
	return db
}
