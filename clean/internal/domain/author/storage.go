package author

import "context"

// Интерфейс "хранилище"
type Storage interface {
	GetOne(ctx context.Context, uuid string) (*Author, error)
	GetAll(ctx context.Context, limit, offset int) ([]*Author, error)
	Create(ctx context.Context, author *Author) (*Author, error)
	Update(ctx context.Context, author *Author) (*Author, error)
	Delete(ctx context.Context, uuid string) error
}
