package teststore

import (
	"go-learn/rest4/internal/app/model"
	"go-learn/rest4/internal/app/store"
)

// Структура "хранилище пользователей"
type UserRepository struct {
	store *Store
	users map[int]*model.User
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
	user.ID = len(r.users) + 1
	r.users[user.ID] = user
	return nil
}

// Поиск пользователя по ID
func (r *UserRepository) Find(id int) (*model.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return user, nil

}

// Поиск пользователя по Email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, store.ErrRecordNotFound
}
