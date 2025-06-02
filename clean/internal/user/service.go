package user

import "context"

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
	return &service{storage: storage}
}

// Получение пользователя по ID
func (s *service) GetUserByUUID(ctx context.Context, uuid string) (*User, error) {
	return s.storage.GetOne(ctx, uuid)
}

// Получение всех пользователей
func (s *service) GetAllUsers(ctx context.Context, limit, offset int) ([]*User, error) {
	return s.storage.GetAll(ctx, limit, offset)
}
