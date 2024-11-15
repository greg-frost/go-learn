package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	"go-learn/base"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

// Путь
var path = base.Dir("goose")

//go:embed migrations/*.sql
var fs embed.FS

func main() {
	fmt.Println(" \n[ GOOSE ]\n ")

	// Подключение к БД
	dsn := "postgres://postgres:admin@localhost:5432/learn?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Выбор каталога
	goose.SetBaseFS(fs)

	// Установка диалекта
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	// Применение миграций
	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Миграции применены!")
}
