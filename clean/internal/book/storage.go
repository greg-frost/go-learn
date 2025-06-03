package book

import "context"

// Интерфейс "хранилище"
type Storage interface {
	GetOne(ctx context.Context, uuid string) (*Book, error)
	GetAll(ctx context.Context, limit, offset int) ([]*Book, error)
	Create(ctx context.Context, book *Book) (*Book, error)
	Update(ctx context.Context, book *Book) (*Book, error)
	Delete(ctx context.Context, book *Book) error
}
