package author

import (
	"context"

	"github.com/google/uuid"
)

// Интерфейс "сервис"
type Service interface {
	GetAuthorByUUID(ctx context.Context, uuid string) (*Author, error)
	GetAllAuthors(ctx context.Context, limit, offset int) ([]*Author, error)
	CreateAuthor(ctx context.Context, dto *CreateAuthorDTO) (*Author, error)
	UpdateAuthor(ctx context.Context, dto *UpdateAuthorDTO) (*Author, error)
	DeleteAuthor(ctx context.Context, dto *DeleteAuthorDTO) error
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

// Получение автора по ID
func (s *service) GetAuthorByUUID(ctx context.Context, uuid string) (*Author, error) {
	return s.storage.GetOne(ctx, uuid)
}

// Получение всех авторов
func (s *service) GetAllAuthors(ctx context.Context, limit, offset int) ([]*Author, error) {
	return s.storage.GetAll(ctx, limit, offset)
}

// Создание автора
func (s *service) CreateAuthor(ctx context.Context, dto *CreateAuthorDTO) (*Author, error) {
	author := &Author{
		UUID:  uuid.NewString(),
		Name:  dto.Name,
		Email: dto.Email,
	}
	return s.storage.Create(ctx, author)
}

// Обновление автора
func (s *service) UpdateAuthor(ctx context.Context, dto *UpdateAuthorDTO) (*Author, error) {
	author := &Author{
		UUID:  dto.UUID,
		Name:  dto.Name,
		Email: dto.Email,
	}
	return s.storage.Update(ctx, author)
}

// Удаление автора
func (s *service) DeleteAuthor(ctx context.Context, dto *DeleteAuthorDTO) error {
	author := &Author{
		UUID: dto.UUID,
	}
	return s.storage.Delete(ctx, author)
}
