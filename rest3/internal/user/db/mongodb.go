package db

import (
	"context"
	"fmt"

	"go-learn/rest3/internal/user"
	"go-learn/rest3/pkg/logger"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Структура "база данных"
type db struct {
	collection *mongo.Collection
	logger     *logger.Logger
}

// Конструктор хранилища
func NewStorage(database *mongo.Database, collection string, logger *logger.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}

// Создание пользователя
func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	d.logger.Debug("Создание пользователя")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("не удалось создать пользователя: %w", err)
	}

	d.logger.Debug("Получение ObjectID пользователя")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		d.logger.Trace(user)
		return "", fmt.Errorf("не удалось получить ObjectID пользователя: %w", err)
	}
	return oid.Hex(), nil
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
