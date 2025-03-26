package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Конструктор клиента
func NewClient(ctx context.Context, host, port, user, pass, db, authDB string) (*mongo.Database, error) {
	// Строка подключения
	var mongodbURL string
	var isAuth bool
	if user == "" && pass == "" {
		mongodbURL = fmt.Sprintf("mongodb://%s:%s", host, port)
	} else {
		isAuth = true
		mongodbURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", user, pass, host, port)
	}

	// Опции
	clientOptions := options.Client().ApplyURI(mongodbURL)
	if isAuth {
		if authDB == "" {
			authDB = db
		}
		// Авторизация
		clientOptions.SetAuth(options.Credential{
			AuthSource: authDB,
			Username:   user,
			Password:   pass,
		})
	}

	// Подключение
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к mongodb: %w", err)
	}

	// Пинг
	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ошибка пинга к mongodb: %w", err)
	}

	return client.Database(db), nil
}
