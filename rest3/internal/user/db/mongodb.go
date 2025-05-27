package db

import (
	"context"
	"errors"
	"fmt"

	"go-learn/rest3/internal/apperror"
	"go-learn/rest3/internal/user"
	"go-learn/rest3/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
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

// Поиск конкретного пользователя
func (d *db) FindOne(ctx context.Context, id string) (user.User, error) {
	var u user.User

	d.logger.Debug("Получение ObjectID пользователя")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("не удалось получить ObjectID пользователя: %w, hex: %s", err, id)
	}

	d.logger.Debug("Поиск пользователя")
	filter := bson.M{"_id": oid}
	result := d.collection.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return u, apperror.ErrNotFound
			// return u, fmt.Errorf("пользователь не найден: %w, id: %s", err, id)
		}
		return u, fmt.Errorf("не удалось получить пользователя: %w, id: %s", err, id)
	}

	d.logger.Debug("Декодирование пользователя")
	if err := result.Decode(&u); err != nil {
		return u, fmt.Errorf("не удалось декодировать пользователя: %w", err)
	}

	return u, nil
}

// Поиск всех пользователей
func (d *db) FindAll(ctx context.Context) ([]user.User, error) {
	var u []user.User

	d.logger.Debug("Поиск пользователей")
	filter := bson.M{}
	cursor, err := d.collection.Find(ctx, filter)
	if err != nil {
		return u, fmt.Errorf("не удалось получить пользователей: %w", err)
	}

	d.logger.Debug("Декодирование пользователей")
	if err := cursor.All(ctx, &u); err != nil {
		return u, fmt.Errorf("не удалось декодировать пользователей: %w", err)
	}
	if cursor.Err() != nil {
		return u, fmt.Errorf("не удалось прочитать пользователей: %w", cursor.Err())
	}

	return u, nil
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
		return "", fmt.Errorf("не удалось получить ObjectID пользователя: %w, oid: %v", err, oid)
	}
	return oid.Hex(), nil
}

// Обновление пользователя
func (d *db) Update(ctx context.Context, user user.User) error {
	d.logger.Debug("Получение ObjectID пользователя")
	oid, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return fmt.Errorf("не удалось получить ObjectID пользователя: %w, hex: %s", err, user.ID)
	}

	d.logger.Debug("Обновление пользователя")
	filter := bson.M{"_id": oid}
	userBytes, err := bson.Marshal(user)
	if err != nil {
		return fmt.Errorf("не удалось сериализовать пользователя: %w, id: %s", err, user.ID)
	}
	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return fmt.Errorf("не удалось десериализовать пользователя: %w, id: %s", err, user.ID)
	}
	delete(updateUserObj, "_id")
	update := bson.M{"set": updateUserObj}
	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("не удалось обновить пользователя: %w, id: %s", err, user.ID)
	}
	if result.MatchedCount == 0 {
		return apperror.ErrNotFound
		// return fmt.Errorf("пользователь не найден: %w, id: %s", err, user.ID)
	}
	d.logger.Tracef("matched: %d, modified: %d", result.MatchedCount, result.ModifiedCount)

	return nil
}

// Удаление пользователя
func (d *db) Delete(ctx context.Context, id string) error {
	d.logger.Debug("Получение ObjectID пользователя")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("не удалось получить ObjectID пользователя: %w, hex: %s", err, id)
	}

	d.logger.Debug("Удаление пользователя")
	filter := bson.M{"_id": oid}
	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("не удалось удалить пользователя: %w, id: %s", err, id)
	}
	if result.DeletedCount == 0 {
		return apperror.ErrNotFound
		// return fmt.Errorf("пользователь не найден: %w, id: %s", err, id)
	}
	d.logger.Tracef("deleted: %d", result.DeletedCount)

	return nil
}
