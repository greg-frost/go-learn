package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Дескриптор БД
var db *sql.DB

// Структура "альбом"
type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	fmt.Println(" \n[ БАЗА ДАННЫХ ]\n ")

	/* Подключение */

	// БД, адрес, логин и пароль
	dbname := "learn"
	addr := "127.0.0.1:3306"
	username := os.Getenv("DB_USER")
	if username == "" {
		username = "root"
	}
	password := os.Getenv("DB_PASS")

	// Вариант 1
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username, password, addr, dbname)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}

	// Конфигурация
	// cfg := mysql.Config{
	// 	User:   username,
	// 	Passwd: password,
	// 	Net:    "tcp",
	// 	Addr:   addr,
	// 	DBName: dbname,
	// 	// AllowNativePasswords: true,
	// }

	// Вариант 2
	// db, err := sql.Open("mysql", cfg.FormatDSN())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Вариант 3
	// connector, err := mysql.NewConnector(&cfg)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// db = sql.OpenDB(connector)

	defer db.Close()

	// Пинг
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Успешное подключение")
	fmt.Println()

	// Удаление старой таблицы
	_, err = db.Exec("DROP TABLE IF EXISTS album")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Старая таблица удалена")

	// Создание новой таблицы
	_, err = db.Exec(`
		CREATE TABLE album (
			id INT AUTO_INCREMENT NOT NULL,
			title VARCHAR(128) NOT NULL,
			artist VARCHAR(255) NOT NULL,
			price DECIMAL(5,2) NOT NULL,
			PRIMARY KEY (id))
		`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Новая таблица создана")

	// Вставка записей
	insert, err := db.Exec(`
		INSERT INTO album
			(title, artist, price)
		VALUES
			('Blue Train', 'John Coltrane', 56.99),
			('Giant Steps', 'John Coltrane', 63.99),
			('Jeru', 'Gerry Mulligan', 17.99),
			('Jeru (remastered)', 'Gerry Mulligan', 19.99),
			('Sarah Vaughan', 'Sarah Vaughan', 34.98)
		`)
	if err != nil {
		log.Fatal(err)
	}
	inserted, _ := insert.RowsAffected()
	fmt.Printf("Добавлены записи (%d)\n", inserted)
	fmt.Println()

	// Убаление таблицы
	_, err = db.Exec("DROP TABLE album")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Таблица удалена")
}
