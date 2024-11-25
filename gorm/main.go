package main

import (
	"fmt"
	"log"
	"time"

	// "gorm.io/driver/postgres"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	// Конфигурация
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	// Подключение к MySQL
	dsn := "root@tcp(localhost:3306)/learn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), config)

	// Подключение к PostgreSQL
	// dsn := "host=localhost user=postgres password=admin dbname=learn sslmode=disable"
	// db, err := gorm.Open(postgres.Open(dsn), config)

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

	// Создание пользователя
	user := User{
		Name:  "Greg Frost",
		Email: "greg-frost@yandex.ru",
		Age:   37,
		Sessions: []Session{
			{Device: "Gregory's PC", Expires: time.Now().Add(72 * time.Hour)},
			{Device: "Greg's iPhone", Expires: time.Now().Add(24 * time.Hour)},
		},
	}
	if res := db.Create(&user); res.Error != nil {
		log.Println(res.Error)
	} else {
		fmt.Println("Новый пользователь добавлен")
		fmt.Println("Сессии пользователя добавлены")
	}
	fmt.Println()

	// Создание нескольких пользователей
	users := []User{
		{Name: "Morozov Grigoriy", Email: "iam@nonexist.com", Age: 30},
		{Name: "Testerov Tester", Email: "fromthe@void.net"},
	}
	if res := db.Create(&users); res.Error != nil {
		log.Println(res.Error)
	} else {
		fmt.Println("Еще несколько пользователей добавлено")
	}
	fmt.Println()

	// Чтение пользователя
	var firstUser User
	db.First(&firstUser)
	// db.Take(&firstUser)
	fmt.Printf("Первый пользователь:\nName: %s, Email: %s, Age: %d, Sessions: %d\n\n",
		firstUser.Name, firstUser.Email, firstUser.Age, len(firstUser.Sessions))

	var allUsers []User
	db.Find(&allUsers)
	fmt.Println("Все пользователи:")
	for i := 0; i < len(allUsers); i++ {
		fmt.Printf("Name: %s, Email: %s, Age: %d\n",
			allUsers[i].Name, allUsers[i].Email, allUsers[i].Age)
	}
	fmt.Println()

	// Удаление пользователя
	db.Delete(&users, 3)
	fmt.Println("Пользователь удален")
	// fmt.Println()

	// Удаление таблиц
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
