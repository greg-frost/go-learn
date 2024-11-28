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
	Profile  Profile   // Has One
	Sessions []Session // Has Many
}

// Структура "сессия"
type Session struct {
	gorm.Model
	UserID  uint
	Device  string    `gorm:"size:100;index"`
	Expires time.Time `gorm:"index:idx_expires_at,sort:desc"`
}

// Структура "профиль"
type Profile struct {
	gorm.Model
	UserID  uint
	Caption string `gorm:"size:100"`
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
	db.AutoMigrate(&User{}, &Session{}, &Profile{})
	fmt.Println("Таблица пользователей создана")
	fmt.Println("Таблица пользовательских профилей создана")
	fmt.Println("Таблица пользовательских сессий создана")
	fmt.Println()

	/* Создание */

	user := User{
		Name: "Gregory Frost",
		Age:  21,
	}

	// Связь "Has One"
	profile := Profile{
		Caption: "Профиль пользователя",
	}
	user.Profile = profile

	// Связь "Has Many"
	sessions := []Session{
		{Device: "Gregory's PC", Expires: time.Now().Add(72 * time.Hour)},
		{Device: "Greg's iPhone", Expires: time.Now().Add(24 * time.Hour)},
	}
	user.Sessions = sessions

	if res := db.Create(&user); res.Error != nil {
		log.Println(res.Error)
	} else {
		fmt.Println("Новый пользователь добавлен")
		fmt.Println("Сессии пользователя добавлены")
	}
	fmt.Println()

	// Несколько
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

	/* Обновление */

	// Полностью
	user.Email = "greg-frost@yandex.ru"
	db.Model(&user).Updates(user)

	// Частично
	db.Model(&user).Updates(map[string]interface{}{
		"Name": "Greg Frost", "Age": 37,
	})

	// Несколько
	db.Model(&User{}).Where("age < 30").
		Updates(map[string]interface{}{"age": 21})

	/* Чтение записи */

	var firstUser User

	// Первый по ключу
	// db.First(&firstUser)

	// Первый попавшийся
	// db.Take(&firstUser)

	// Связанные данные
	db.Preload("Profile").First(&firstUser)

	fmt.Printf("Первый пользователь:\nName: %s, Email: %s, Age: %d\nProfile: %s\n\n",
		firstUser.Name, firstUser.Email, firstUser.Age, firstUser.Profile.Caption)

	/* Чтение всех записей */

	var allUsers []User

	// Только записи
	db.Find(&allUsers)

	// Связанные данные
	// db.Preload("Sessions").Find(&allUsers)

	fmt.Println("Все пользователи:")
	for i := 0; i < len(allUsers); i++ {
		fmt.Printf("Name: %s, Email: %s, Age: %d\n",
			allUsers[i].Name, allUsers[i].Email, allUsers[i].Age)
	}
	fmt.Println()

	// Пагинация
	var pageUsers []User
	page, limit := 2, 1
	db.Order("age desc").Offset((page - 1) * limit).Limit(limit).Find(&pageUsers)
	fmt.Printf("Средний пользователь (по возрасту):\nName: %s, Email: %s, Age: %d\n\n",
		pageUsers[0].Name, pageUsers[0].Email, pageUsers[0].Age)

	// Удаление записи
	db.Delete(&users, 3)
	fmt.Println("Пользователь удален")
	fmt.Println()

	/* Транзакции */

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		fmt.Println("Транзакция отменена")
	} else {
		tx.Commit()
		fmt.Println("Транзакция выполнена")
	}
	fmt.Println()

	// Удаление таблиц
	db.Exec("DROP TABLE users")
	db.Exec("DROP TABLE sessions")
	db.Exec("DROP TABLE profiles")
	fmt.Println("Таблицы удалены")
}
