package apiserver

import (
	"database/sql"
	"io"
	"net/http"

	"go-learn/rest4/internal/app/store"
	"go-learn/rest4/internal/app/store/sqlstore"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Структура "сервер API"
type APIServer struct {
	config *Config
	logger *logrus.Logger
	store  store.Store
	router *mux.Router
}

// Конструктор сервера
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Запуск сервера
func (s *APIServer) Start() error {
	s.logger.Info("Конфигурирование логгера")
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("Конфигурирование хранилища")
	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Конфигурирование роутера")
	s.configureRouter()

	// Запуск сервера
	s.logger.Info("Запуск сервера API")
	s.logger.Info("Ожидаю соединений...")
	s.logger.Infof("(на http://%s)", s.config.BindAddr)

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// Конфигурирование логгера
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil
}

// Конфигурирование хранилища
func (s *APIServer) configureStore() error {
	db, err := sql.Open("postgres", s.config.Store.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.store = sqlstore.New(db)

	return nil
}

// Конфигурирование роутера
func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

// Обработчик приветствия
func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Привет, мир!")
	}
}
