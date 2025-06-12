package book

import (
	"context"

	"github.com/google/uuid"
)

// Структура "сервис"
type service struct {
	storage Storage
}

// Конструктор сервиса
func NewService(storage Storage) *service {
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
		UUID:  uuid.NewString(),
		Title: dto.Title,
		// Author: dto.AuthorUUID,
		Year: dto.Year,
	}
	return s.storage.Create(ctx, book)
}

// Обновление книги
func (s *service) UpdateBook(ctx context.Context, dto *UpdateBookDTO) (*Book, error) {
	book := &Book{
		UUID:  dto.UUID,
		Title: dto.Title,
		// Author:    dto.AuthorUUID,
		Year:      dto.Year,
		Busy:      dto.Busy,
		OwnerUUID: dto.OwnerUUID,
	}
	return s.storage.Update(ctx, book)
}

// Удаление книги
func (s *service) DeleteBook(ctx context.Context, dto *DeleteBookDTO) error {
	book := &Book{
		UUID: dto.UUID,
	}
	return s.storage.Delete(ctx, book)
}
