package store

import "go-learn/rest4/internal/app/model"

// Интерфейс "хранилище пользователей"
type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}
