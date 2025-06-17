package book

import (
	"context"
	"errors"

	"go-learn/clean/internal/domain/book"
	"go-learn/clean/pkg/client/memory"
)

// Структура "хранилище"
type storage struct {
	client *memory.Client
}

// Конструктор хранилища
func NewStorage() book.Storage {
	return &storage{
		client: memory.NewClient(),
	}
}

// Получение конкретной книги
func (s *storage) GetOne(ctx context.Context, uuid string) (*book.Book, error) {
	value, ok := s.client.Get(uuid)
	if !ok {
		return nil, errors.New("не найдено")
	}
	return value.(*book.Book), nil
}

// Получение всех книг
func (s *storage) GetAll(ctx context.Context, limit, offset int) ([]*book.Book, error) {
	books := make([]*book.Book, 0, limit)
	for _, value := range s.client.Values() {
		books = append(books, value.(*book.Book))
	}
	if offset > len(books) {
		offset = len(books)
	}
	books = books[offset:]
	if len(books) > limit {
		books = books[:limit]
	}
	return books, nil
}

// Создание книги
func (s *storage) Create(ctx context.Context, book *book.Book) (*book.Book, error) {
	if ok := s.client.Put(book.UUID, book); !ok {
		return nil, errors.New("не сохранено")
	}
	return book, nil
}

// Обновление книги
func (s *storage) Update(ctx context.Context, book *book.Book) (*book.Book, error) {
	if _, ok := s.client.Get(book.UUID); !ok {
		return nil, errors.New("не найдено")
	}
	if ok := s.client.Put(book.UUID, book); !ok {
		return nil, errors.New("не обновлено")
	}
	return book, nil
}

// Удаление книги
func (s *storage) Delete(ctx context.Context, uuid string) error {
	if ok := s.client.Delete(uuid); !ok {
		return errors.New("не удалено")
	}
	return nil
}
