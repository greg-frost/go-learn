package main

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log"
	"path/filepath"

	"go-learn/base"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

// Структура "мигратор"
type Migrator struct {
	driver source.Driver
}

// Обязательное получение нового мигратора
func MustNewMigrator(files embed.FS, dir string) *Migrator {
	d, err := iofs.New(files, dir)
	if err != nil {
		panic(err)
	}
	return &Migrator{driver: d}
}

// Применение миграций
func (m *Migrator) ApplyMigrations(db *sql.DB) error {
	// Получение драйвера БД
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("не удалось создать экземпляр БД: %v", err)
	}

	// Получение мигратора
	migrator, err := migrate.NewWithInstance(
		"migration_embeded_sql_files",
		m.driver, "psql_db", driver,
	)
	if err != nil {
		return fmt.Errorf("не удалось создать экземпляр мигратора: %v", err)
	}
	defer migrator.Close()

	// Миграции
	if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("не удалось применить миграцию: %v", err)
	}

	return nil
}

// Путь
var path = base.Dir("migrate")

//go:embed migrations/*.sql
var fs embed.FS

func main() {
	fmt.Println(" \n[ GO-MIGRATE ]\n ")

	// Создание мигратора
	migrator := MustNewMigrator(fs, filepath.Join(path, "migrations"))

	// Подключение к БД
	dsn := "postgres://postgres:admin@localhost:5432/learn?sslmode=disable"
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Применение миграций
	err = migrator.ApplyMigrations(conn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Миграции применены!")
}
