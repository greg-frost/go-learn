package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Структура "пользователь"
type User struct {
	Id   int
	Name string
}

func main() {
	fmt.Println(" \n[ MYSQL ]\n ")

	// Подключение
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/golearn")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Пинг
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Успешное подключение")

	// Очистка
	truncate, err := db.Query("TRUNCATE TABLE users")
	if err != nil {
		log.Fatal(err)
	}
	defer truncate.Close()
	fmt.Println("Таблица очищена")

	// Вставка
	insert, err := db.Query("INSERT INTO users(name) VALUES('go'), ('greg'), ('frost')")
	if err != nil {
		log.Fatal(err)
	}
	defer insert.Close()
	fmt.Println("Созданы записи")

	// Чтение
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer results.Close()

	// Вывод
	for results.Next() {
		var user User
		err = results.Scan(&user.Id, &user.Name)
		fmt.Println(user)
	}
}
