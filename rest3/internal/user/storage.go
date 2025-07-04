package user

import "context"

// Интерфейс "хранилище"
type Storage interface {
	FindOne(ctx context.Context, id string) (User, error)
	FindAll(ctx context.Context) ([]User, error)
	Create(ctx context.Context, user User) (string, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
}
