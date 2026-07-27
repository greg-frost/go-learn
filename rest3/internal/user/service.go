package user

import (
	"context"

	"go-learn/rest3/pkg/logger"
)

// Интерфейс "сервис"
type Service interface {
	FindOne(ctx context.Context, id string) (User, error)
	FindAll(ctx context.Context) ([]User, error)
	Create(ctx context.Context, dto CreateUserDTO) (User, error)
	Update(ctx context.Context, dto UpdateUserDTO) error
	Delete(ctx context.Context, id string) error
}

// Структура "сервис"
type service struct {
	storage Storage
	logger  *logger.Logger
}

// Конструктор сервиса
func NewService(storage Storage, logger *logger.Logger) Service {
	return &service{
		storage: storage,
		logger:  logger,
	}
}

// Поиск конкретного пользователя
func (s *service) FindOne(ctx context.Context, id string) (User, error) {
	s.logger.Debug("Поиск конкретного пользователя")
	s.logger.Tracef("id: %s", id)
	return s.storage.FindOne(ctx, id)
}

// Поиск всех пользователей
func (s *service) FindAll(ctx context.Context) ([]User, error) {
	s.logger.Debug("Поиск всех пользователей")
	return s.storage.FindAll(ctx)
}

// Создание пользователя
func (s *service) Create(ctx context.Context, dto CreateUserDTO) (User, error) {
	s.logger.Debug("Создание пользователя")
	s.logger.Trace(dto)
	panic("не реализовано")
	// return s.storage.Create(ctx, dto)
}

// Обновление пользователя
func (s *service) Update(ctx context.Context, dto UpdateUserDTO) error {
	s.logger.Debug("Обновление пользователя")
	s.logger.Trace(dto)
	panic("не реализовано")
	// return s.storage.Update(ctx, dto)
}

// Удаление пользователя
func (s *service) Delete(ctx context.Context, id string) error {
	s.logger.Debug("Удаление пользователя")
	s.logger.Tracef("id: %s", id)
	return s.storage.Delete(ctx, id)
}
