package book

import (
	"context"

	"go-learn/clean/internal/domain/book"
)

// Интерфейс "сервис"
type Service interface {
	GetBookByUUID(ctx context.Context, uuid string) (*book.Book, error)
	GetAllBooks(ctx context.Context, limit, offset int) ([]*book.Book, error)
	CreateBook(ctx context.Context, dto *book.CreateBookDTO) (*book.Book, error)
}
