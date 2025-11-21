package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Структура "хранилище"
type Store struct {
	config *Config
	db     *sql.DB
}

// Конструктор хранилища
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Открытие соединения
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db

	return nil
}

// Закрытие соединения
func (s *Store) Close() {
	s.db.Close()
}
