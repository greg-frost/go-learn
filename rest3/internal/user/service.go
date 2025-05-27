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

// Создание пользователя
func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (User, error) {
	panic("не реализовано")
}
