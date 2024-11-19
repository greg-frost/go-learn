package main

import (
	"fmt"
	"log"

	// "gorm.io/driver/postgres"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Структура "пользователь"
type User struct {
	gorm.Model
	Name  string `gorm:"size:50"`
	Email string `gorm:"type:varchar(100);uniqueIndex"`
	Age   int    `gorm:"default:18"`
}

func main() {
	fmt.Println(" \n[ GORM ]\n ")

	/* Подключение к БД */

	// MySQL
	dsn := "root@tcp(localhost:3306)/learn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// PostgreSQL
	// dsn := "host=localhost user=postgres password=admin dbname=learn sslmode=disable"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Соединение установлено!")

	/* Миграции */

	// Таблица пользователей
	db.AutoMigrate(&User{})
	fmt.Println("Таблица пользователей создана.")
}
