package user

import (
	"context"

	"github.com/google/uuid"
)

// Интерфейс "сервис"
type Service interface {
	GetUserByUUID(ctx context.Context, uuid string) (*User, error)
	GetAllUsers(ctx context.Context, limit, offset int) ([]*User, error)
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

// Получение пользователя по ID
func (s *service) GetUserByUUID(ctx context.Context, uuid string) (*User, error) {
	return s.storage.GetOne(ctx, uuid)
}

// Получение всех пользователей
func (s *service) GetAllUsers(ctx context.Context, limit, offset int) ([]*User, error) {
	return s.storage.GetAll(ctx, limit, offset)
}

// Создание пользователя
func (s *service) CreateUser(ctx context.Context, dto *CreateUserDTO) (*User, error) {
	user := &User{
		UUID:    uuid.NewString(),
		Name:    dto.Name,
		Address: dto.Address,
		Email:   dto.Email,
		Age:     dto.Age,
	}
	return s.storage.Create(ctx, user)
}

// Обновление пользователя
func (s *service) UpdateUser(ctx context.Context, dto *UpdateUserDTO) (*User, error) {
	user := &User{
		UUID:    dto.UUID,
		Name:    dto.Name,
		Address: dto.Address,
		Email:   dto.Email,
		Age:     dto.Age,
	}
	return s.storage.Update(ctx, user)
}
