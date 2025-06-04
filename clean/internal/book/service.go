package book

import (
	"context"

	"github.com/google/uuid"
)

// Интерфейс "сервис"
type Service interface {
	GetBookByUUID(ctx context.Context, uuid string) (*Book, error)
	GetAllBooks(ctx context.Context, limit, offset int) ([]*Book, error)
	CreateBook(ctx context.Context, dto *CreateBookDTO) (*Book, error)
}

// Структура "сервис"
type service struct {
	storage Storage
}

// Конструктор сервиса
func NewService(storage Storage) Service {
	return &service{
		storage: storage,
	}
}

// Получение книги по ID
func (s *service) GetBookByUUID(ctx context.Context, uuid string) (*Book, error) {
	return s.storage.GetOne(ctx, uuid)
}

// Получение всех книг
func (s *service) GetAllBooks(ctx context.Context, limit, offset int) ([]*Book, error) {
	return s.storage.GetAll(ctx, limit, offset)
}

// Создание книги
func (s *service) CreateBook(ctx context.Context, dto *CreateBookDTO) (*Book, error) {
	book := &Book{
		UUID:   uuid.NewString(),
		Title:  dto.Title,
		Author: dto.Author,
		Year:   dto.Year,
	}
	return s.storage.Create(ctx, book)
}

// Обновление книги
func (s *service) UpdateBook(ctx context.Context, dto *UpdateBookDTO) (*Book, error) {
	book := &Book{
		UUID:   dto.UUID,
		Title:  dto.Title,
		Author: dto.Author,
		Year:   dto.Year,
		Busy:   dto.Busy,
		User:   dto.User,
	}
	return s.storage.Update(ctx, book)
}
