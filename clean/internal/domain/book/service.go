package book

import (
	"context"
	"errors"

	"go-learn/clean/internal/domain/author"

	"github.com/google/uuid"
)

// Структура "сервис"
type service struct {
	storage       Storage
	authorService author.Service
}

// Конструктор сервиса
// Замечание:
// Извлечение интерфейса к месту использования не позволяет здесь возвращать интерфейс,
// т.к. это создает циклическую зависимость, да и не обязан сервис знать своих потребитетелей
// (можно было бы разделить пакеты модели и сервиса, что решило бы проблему циклического импорта)
func NewService(storage Storage, aservice author.Service) *service {
	return &service{
		storage:       storage,
		authorService: aservice,
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
	author, err := s.authorService.GetAuthorByUUID(ctx, dto.AuthorUUID)
	if err != nil {
		return nil, errors.New("автор не найден")
	}
	book := &Book{
		UUID:   uuid.NewString(),
		Title:  dto.Title,
		Author: *author,
		Year:   dto.Year,
	}
	return s.storage.Create(ctx, book)
}

// Обновление книги
func (s *service) UpdateBook(ctx context.Context, dto *UpdateBookDTO) (*Book, error) {
	author, err := s.authorService.GetAuthorByUUID(ctx, dto.AuthorUUID)
	if err != nil {
		return nil, errors.New("автор не найден")
	}
	book := &Book{
		UUID:      dto.UUID,
		Title:     dto.Title,
		Author:    *author,
		Year:      dto.Year,
		Busy:      dto.Busy,
		OwnerUUID: dto.OwnerUUID,
	}
	return s.storage.Update(ctx, book)
}

// Удаление книги
func (s *service) DeleteBook(ctx context.Context, uuid string) error {
	return s.storage.Delete(ctx, uuid)
}
