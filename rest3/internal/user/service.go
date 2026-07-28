package user

import (
	"context"

	"go-learn/rest3/pkg/logger"
)

// Интерфейс "сервис"
type Service interface {
	GetUserByID(ctx context.Context, id string) (User, error)
	GetUsers(ctx context.Context) ([]User, error)
	CreateUser(ctx context.Context, dto CreateUserDTO) (User, error)
	UpdateUser(ctx context.Context, dto UpdateUserDTO) error
	DeleteUser(ctx context.Context, id string) error
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

// Получение конкретного пользователя
func (s *service) GetUserByID(ctx context.Context, id string) (User, error) {
	s.logger.Debug("Получение конкретного пользователя")
	s.logger.Tracef("id: %s", id)
	return s.storage.FindOne(ctx, id)
}

// Получение всех пользователей
func (s *service) GetUsers(ctx context.Context) ([]User, error) {
	s.logger.Debug("Получение всех пользователей")
	return s.storage.FindAll(ctx)
}

// Создание пользователя
func (s *service) CreateUser(ctx context.Context, dto CreateUserDTO) (User, error) {
	s.logger.Debug("Создание пользователя")
	user := User{
		Email:    dto.Email,
		Username: dto.Username,
	}
	err := user.encryptPassword(dto.Password)
	if err != nil {
		return User{}, err
	}
	dto.Password = ""
	s.logger.Trace(dto)

	return s.storage.Create(ctx, user)
}

// Обновление пользователя
func (s *service) UpdateUser(ctx context.Context, dto UpdateUserDTO) error {
	s.logger.Debug("Обновление пользователя")
	user := User{
		ID:       dto.ID,
		Email:    dto.Email,
		Username: dto.Username,
	}
	err := user.encryptPassword(dto.Password)
	if err != nil {
		return err
	}
	dto.Password = ""
	s.logger.Trace(dto)

	return s.storage.Update(ctx, user)
}

// Удаление пользователя
func (s *service) DeleteUser(ctx context.Context, id string) error {
	s.logger.Debug("Удаление пользователя")
	s.logger.Tracef("id: %s", id)
	return s.storage.Delete(ctx, id)
}
