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
	_ "github.com/lib/pq"
)

// Путь
var path = base.Dir("migrate")

// Миграции
var migrations = filepath.Join(path, "migrations")

//go:embed migrations/*.sql
var fs embed.FS

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

	// Миграция вверх (до конца)
	if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("не удалось применить миграцию: %v", err)
	}

	return nil
}

// Откат миграций
func (m *Migrator) RevertMigrations(db *sql.DB) error {
	// Получение драйвера БД
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("не удалось создать экземпляр БД: %v", err)
	}

	// Получение мигратора (через БД)
	migrator, err := migrate.NewWithDatabaseInstance(
		"file:///"+migrations,
		"postgres", driver,
	)
	if err != nil {
		return fmt.Errorf("не удалось создать экземпляр мигратора: %v", err)
	}
	defer migrator.Close()

	// Миграция вниз (до конца)
	if err := migrator.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("не удалось откатить миграцию: %v", err)
	}

	return nil
}

func main() {
	fmt.Println(" \n[ GO-MIGRATE ]\n ")

	// Создание мигратора
	migrator := MustNewMigrator(fs, migrations)

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
	fmt.Println("Миграции применены!")

	// Откат миграций
	err = migrator.RevertMigrations(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Миграции отменены...")
}
