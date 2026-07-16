package store

import "go-learn/rest4/internal/app/model"

// Интерфейс "хранилище пользователей"
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
