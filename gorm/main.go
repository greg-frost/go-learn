package main

import (
	"fmt"
	"log"
	"time"

	// "gorm.io/driver/postgres"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Структура "пользователь"
type User struct {
	gorm.Model
	Name     string    `gorm:"size:50"`
	Email    string    `gorm:"type:varchar(100);unique"`
	Age      int32     `gorm:"default:18"`
	Sessions []Session `gorm:"foreignKey:UserID"`
}

// Структура "сессия"
type Session struct {
	gorm.Model
	UserID  uint
	Name    string    `gorm:"size:100;index"`
	Expires time.Time `gorm:"index:idx_expires_at,sort:desc"`
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
	fmt.Println()

	/* Миграции */

	// Пользователи и их сессии
	db.AutoMigrate(&User{}, &Session{})

	fmt.Println("Таблица пользователей создана.")
	fmt.Println("Таблица пользовательских сессий создана.")
}
