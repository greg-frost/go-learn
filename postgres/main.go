package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// Структура "пользователь"
type User struct {
	Id   int
	Name string
}

func main() {
	fmt.Println(" \n[ POSTGRESQL ]\n ")

	start := time.Now()

	// Подключение
	conn := "user=postgres password=admin dbname=golearn sslmode=disable"
	db, err := sql.Open("postgres", conn)
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
	fmt.Println()

	// Создание таблицы
	create, err := db.Query(`
			CREATE TABLE IF NOT EXISTS users (
				id SERIAL PRIMARY KEY,
				name VARCHAR(50) NOT NULL
			)
		`)
	if err != nil {
		log.Fatal(err)
	}
	defer create.Close()
	fmt.Println("Таблица создана")

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
	fmt.Println()

	var user User

	// Все записи
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer results.Close()
	fmt.Println("Все:")
	for results.Next() {
		err = results.Scan(&user.Id, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id = %d, name = %s\n", user.Id, user.Name)
	}
	fmt.Println()

	// Первая запись
	result, err := db.Query("SELECT name FROM users WHERE id=$1", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()
	fmt.Println("Первая:")
	if result.Next() {
		err = result.Scan(&user.Name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("name = %s\n", user.Name)
	}
	fmt.Println()

	// Убаление таблицы
	drop, err := db.Query("DROP TABLE users")
	if err != nil {
		log.Fatal(err)
	}
	defer drop.Close()
	fmt.Println("Таблица удалена")

	end := time.Now()
	fmt.Println("Время выполнения:", end.Sub(start))
}
