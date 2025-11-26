package teststore

import (
	"fmt"
	"go-learn/rest4/internal/app/model"
)

// Структура "хранилище пользователей"
type UserRepository struct {
	store *Store
	users map[string]*model.User
}

// Создание пользователя
func (r *UserRepository) Create(user *model.User) error {
	// Валидация
	if err := user.Validate(); err != nil {
		return err
	}

	// Подготовка
	if err := user.BeforeCreate(); err != nil {
		return err
	}

	// Сохранение
	r.users[user.Email] = user
	user.ID = len(r.users)

	return nil
}

// Поиск пользователя по Email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	user, ok := r.users[email]
	if !ok {
		return nil, fmt.Errorf("пользователь %s не найден", email)
	}

	return user, nil
}
