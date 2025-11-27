package apiserver

import (
	"database/sql"
	"net/http"

	"go-learn/rest4/internal/app/store/sqlstore"
)

// Запуск сервера
func Start(config *Config) error {
	// База данных
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	// Хранилище и сервер
	store := sqlstore.New(db)
	s := NewServer(store)

	// Логгер
	if err := s.configureLogger(config.LogLevel); err != nil {
		return err
	}

	// Запуск сервера
	s.logger.Info("Запуск сервера API")
	s.logger.Info("Ожидаю соединений...")
	s.logger.Infof("(на http://%s)", config.BindAddr)

	return http.ListenAndServe(config.BindAddr, s)
}

// Конструктор БД
func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
