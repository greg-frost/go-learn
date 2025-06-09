package user

import (
	"context"
	"errors"

	"go-learn/clean/internal/domain/user"
	"go-learn/clean/pkg/client/memory"
)

// Структура "хранилище"
type storage struct {
	client *memory.Client
}

// Конструктор хранилища
func NewStorage() user.Storage {
	return &storage{
		client: memory.NewClient(),
	}
}

// Получение конкретного пользователя
func (s *storage) GetOne(ctx context.Context, uuid string) (*user.User, error) {
	value, ok := s.client.Get(uuid)
	if !ok {
		return nil, errors.New("не найдено")
	}
	return value.(*user.User), nil
}

// Получение всех пользователей
func (s *storage) GetAll(ctx context.Context, limit, offset int) ([]*user.User, error) {
	users := make([]*user.User, 0, limit)
	for _, value := range s.client.Values() {
		users = append(users, value.(*user.User))
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
	if ok := s.client.Put(user.UUID, user); !ok {
		return nil, errors.New("не сохранено")
	}
	return user, nil
}

// Обновление пользователя
func (s *storage) Update(ctx context.Context, user *user.User) (*user.User, error) {
	if _, ok := s.client.Get(user.UUID); !ok {
		return nil, errors.New("не найдено")
	}
	if ok := s.client.Put(user.UUID, user); !ok {
		return nil, errors.New("не обновлено")
	}
	return user, nil
}

// Удаление пользователя
func (s *storage) Delete(ctx context.Context, user *user.User) error {
	if ok := s.client.Delete(user.UUID); !ok {
		return errors.New("не удалено")
	}
	return nil
}
