package user

import "context"

// Интерфейс "хранилище"
type Storage interface {
	GetOne(ctx context.Context, uuid string) (*User, error)
	GetAll(ctx context.Context, limit, offset int) ([]*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, user *User) error
}
