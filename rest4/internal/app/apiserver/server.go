package apiserver

import (
	"go-learn/rest4/internal/app/store"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Структура "сервер"
type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

// Конструктор сервера
func NewServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}

// Обработка запроса
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Конфигурирование логгера
func (s *server) configureLogger(logLevel string) error {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil
}

// Конфигурирование роутера
func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.handleUsersCreate()).
		Methods(http.MethodPost)
}

// Обработчик создания пользователя
func (s *server) handleUsersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Привет, пользователи!")
	}
}
