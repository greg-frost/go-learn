package book

import (
	"context"

	"go-learn/clean/internal/domain/book"
)

// Интерфейс "сервис"
// Замечание:
// Перенос интерфейса сюда, т.е. ближе к месту использования,
// не позволяет возвращать интерфейс в самом сервисе,
// т.к. это создает циклическую зависимость
type Service interface {
	GetBookByUUID(ctx context.Context, uuid string) (*book.Book, error)
	GetAllBooks(ctx context.Context, limit, offset int) ([]*book.Book, error)
	CreateBook(ctx context.Context, dto *book.CreateBookDTO) (*book.Book, error)
	UpdateBook(ctx context.Context, dto *book.UpdateBookDTO) (*book.Book, error)
	DeleteBook(ctx context.Context, uuid string) error
}
