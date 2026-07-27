package user

import (
	"context"

	"go-learn/rest3/pkg/logger"
)

// Структура "сервис"
type Service struct {
	storage Storage
	logger  *logger.Logger
}

// Конструктор сервиса
func NewService(storage Storage, logger *logger.Logger) *Service {
	return &Service{
		storage: storage,
		logger:  logger,
	}
}

// Создание пользователя
func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (User, error) {
	panic("не реализовано")
}
