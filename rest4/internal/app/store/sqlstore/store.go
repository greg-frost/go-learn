package sqlstore

import (
	"database/sql"

	"go-learn/rest4/internal/app/store"

	_ "github.com/lib/pq"
)

// Структура "хранилище"
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// Конструктор хранилища
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Получение хранилища пользователей
func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
		}
	}
	return s.userRepository
}
