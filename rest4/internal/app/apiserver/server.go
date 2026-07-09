package apiserver

import (
	"encoding/json"
	"net/http"

	"go-learn/rest4/internal/app/model"
	"go-learn/rest4/internal/app/store"

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
func newServer(store store.Store) *server {
	// Сервер
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	// Роутер
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
	// Структура "запрос"
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Обработка запроса
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(request)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

// Обработка ошибки
func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

// Обработка ответа
func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
