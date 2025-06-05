package user

import (
	"context"
	"errors"

	"go-learn/clean/internal/domain/user"
)

// Структура "хранилище"
type storage struct {
	m map[string]*user.User
}

// Конструктор хранилища
func NewStorage() user.Storage {
	return &storage{
		m: make(map[string]*user.User),
	}
}

// Получение конкретного пользователя
func (s *storage) GetOne(ctx context.Context, uuid string) (*user.User, error) {
	user, ok := s.m[uuid]
	if !ok {
		return nil, errors.New("не найдено")
	}
	return user, nil
}

// Получение всех пользователей
func (s *storage) GetAll(ctx context.Context, limit, offset int) ([]*user.User, error) {
	users := make([]*user.User, 0, limit)
	for _, user := range s.m {
		users = append(users, user)
	}
	if offset > len(users) {
		offset = len(users)
	}
	users = users[offset:]
	if len(users) > limit {
		users = users[:limit]
	}
	return users, nil
}

// Создание пользователя
func (s *storage) Create(ctx context.Context, user *user.User) (*user.User, error) {
	s.m[user.UUID] = user
	return user, nil
}

// Обновление пользователя
func (s *storage) Update(ctx context.Context, user *user.User) (*user.User, error) {
	if _, ok := s.m[user.UUID]; !ok {
		return nil, errors.New("не найдено")
	}
	s.m[user.UUID] = user
	return user, nil
}

// Удаление пользователя
func (s *storage) Delete(ctx context.Context, user *user.User) error {
	if _, ok := s.m[user.UUID]; !ok {
		return errors.New("не найдено")
	}
	delete(s.m, user.UUID)
	return nil
}
