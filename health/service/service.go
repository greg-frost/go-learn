package service

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

// Проверка подключения к БД
func Ping(ctx context.Context) error {
	// Коннект
	conn := "user=postgres password=admin dbname=learn sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return err
	}
	defer db.Close()

	// Пинг
	err = db.Ping()
	if err != nil {
		return err
	}

	// Таймаут
	if ctx.Err() != nil {
		return ctx.Err()
	}

	return nil
}
