package book

import "context"

// Интерфейс "сервис"
type Service interface {
	GetBookByUUID(ctx context.Context, uuid string) (*Book, error)
	GetAllBooks(ctx context.Context, limit, offset int) ([]*Book, error)
	TakeBook(ctx context.Context, book *Book, user string) error
}

// Структура "сервис"
type service struct {
	storage Storage
}

// Конструктор сервиса
func NewService(storage Storage) Service {
	return &service{storage: storage}
}

// Получение книги по ID
func (s *service) GetBookByUUID(ctx context.Context, uuid string) (*Book, error) {
	return s.storage.GetOne(ctx, uuid)
}

// Получение всех книг
func (s *service) GetAllBooks(ctx context.Context, limit, offset int) ([]*Book, error) {
	return s.storage.GetAll(ctx, limit, offset)
}

// Взять книгу
func (s *service) TakeBook(ctx context.Context, book *Book, user string) error {
	return book.Take(user)
}
