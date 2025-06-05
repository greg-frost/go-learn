package book

import (
	"context"
	"errors"

	"go-learn/clean/internal/domain/book"
)

// Структура "хранилище"
type storage struct {
	m map[string]*book.Book
}

// Конструктор хранилища
func NewStorage() book.Storage {
	return &storage{
		m: make(map[string]*book.Book),
	}
}

// Получение конкретной книги
func (s *storage) GetOne(ctx context.Context, uuid string) (*book.Book, error) {
	book, ok := s.m[uuid]
	if !ok {
		return nil, errors.New("не найдено")
	}
	return book, nil
}

// Получение всех книг
func (s *storage) GetAll(ctx context.Context, limit, offset int) ([]*book.Book, error) {
	books := make([]*book.Book, 0, limit)
	for _, book := range s.m {
		books = append(books, book)
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
	s.m[book.UUID] = book
	return book, nil
}

// Обновление книги
func (s *storage) Update(ctx context.Context, book *book.Book) (*book.Book, error) {
	if _, ok := s.m[book.UUID]; !ok {
		return nil, errors.New("не найдено")
	}
	s.m[book.UUID] = book
	return book, nil
}

// Удаление книги
func (s *storage) Delete(ctx context.Context, book *book.Book) error {
	if _, ok := s.m[book.UUID]; !ok {
		return errors.New("не найдено")
	}
	delete(s.m, book.UUID)
	return nil
}
