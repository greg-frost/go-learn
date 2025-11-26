package teststore

import (
	"go-learn/rest4/internal/app/model"
	"go-learn/rest4/internal/app/store"
)

// Структура "хранилище"
type Store struct {
	userRepository *UserRepository
}

// Конструктор хранилища
func New() *Store {
	return &Store{}
}

// Получение хранилища пользователей
func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
			users: make(map[string]*model.User),
		}
	}
	return s.userRepository
}
