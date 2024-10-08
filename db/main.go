package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println(" \n[ БАЗЫ ДАННЫХ ]\n ")

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
	// db := sql.OpenDB(connector)

	defer db.Close()

	// Пинг
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Успешное подключение")
}
