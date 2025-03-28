package db

import (
	"context"

	"go-learn/rest3/internal/user"
	"go-learn/rest3/pkg/logger"
)

// Структура "база данных"
type db struct {
}

// Конструктор хранилища
func NewStogare(collection string, logger *logger.Logger) user.Storage {
	return &db{}
}

// Создание пользователя
func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	return "", nil
}

// Поиск конкретного пользователя
func (d *db) FindOne(ctx context.Context, id string) (user.User, error) {
	return user.User{}, nil
}

// Обновление пользователя
func (d *db) Update(ctx context.Context, user user.User) error {
	return nil
}

// Удаление пользователя
func (d *db) Delete(ctx context.Context, id string) error {
	return nil
}
