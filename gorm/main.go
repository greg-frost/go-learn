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
	Device  string    `gorm:"size:100;index"`
	Expires time.Time `gorm:"index:idx_expires_at,sort:desc"`
}

func main() {
	fmt.Println(" \n[ GORM ]\n ")

	// Подключение к MySQL
	dsn := "root@tcp(localhost:3306)/learn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Подключение к PostgreSQL
	// dsn := "host=localhost user=postgres password=admin dbname=learn sslmode=disable"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Соединение установлено!")
	fmt.Println()

	// Создание таблиц (миграции)
	db.AutoMigrate(&User{}, &Session{})
	fmt.Println("Таблица пользователей создана")
	fmt.Println("Таблица пользовательских сессий создана")
	fmt.Println()

	// Создание записей
	user := User{
		Name: "Greg Frost",
		Age:  37,
		Sessions: []Session{
			{Device: "Gregory's PC", Expires: time.Now().Add(72 * time.Hour)},
			{Device: "Greg's iPhone", Expires: time.Now().Add(24 * time.Hour)},
		},
	}
	result := db.Create(&user)
	if result.Error != nil {
		log.Println(err)
	}
	fmt.Println("Новый пользователь добавлен")
	fmt.Println("Сессии пользователя добавлены")
	fmt.Println()

	// Удаление записей
	db.Delete(&user, 1)
	fmt.Println("Пользователь удален")
	// fmt.Println()

	// Убаление таблиц
	// db.Exec("DROP TABLE users")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// db.Exec("DROP TABLE sessions")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Таблицы удалены")
}
