package book

import "context"

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
		Title:  dto.Title,
		Author: dto.Author,
		Year:   dto.Year,
	}
	return s.storage.Create(ctx, book)
}
