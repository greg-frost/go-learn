package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"go-learn/rest4/internal/app/model"
	"go-learn/rest4/internal/app/store"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

const (
	sessionName        = "gopherschool" // Имя сессии
	ctxKeyUser  ctxKey = iota           // Ключ контекста
)

// Ошибки
var (
	errIncorrectEmailOrPassword = errors.New("неверный email или пароль")
	errNotAuthenticated         = errors.New("пользователь не аутентифицирован")
)

// Тип "ключ контекста"
type ctxKey int8

// Структура "сервер"
type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

// Конструктор сервера
func newServer(store store.Store, sessionStore sessions.Store) *server {
	// Сервер
	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,
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
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods(http.MethodPost)
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods(http.MethodPost)

	// Только для /private/***
	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authenticateUser)
	private.HandleFunc("/whoami", s.handleWhoami()).Methods(http.MethodGet)
}

// Аутентификация пользователя
func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получение сессии
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		// Извлечение ID пользователя
		id, ok := session.Values["user_id"]
		if !ok {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		// Поиск пользователя
		u, err := s.store.User().Find(id.(int))
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		// Сохранение пользователя в контексте
		ctx := context.WithValue(r.Context(), ctxKeyUser, u)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r) // Вызов следующего обработчика
	})
}

// Обработчик показа пользователя
func (s *server) handleWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK,
			r.Context().Value(ctxKeyUser).(*model.User))
	}
}

// Обработчик создания пользователя
func (s *server) handleUsersCreate() http.HandlerFunc {
	// Структура "запрос"
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Парсинг запроса
		req := new(request)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		// Создание пользователя
		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize() // Очистка данных

		s.respond(w, r, http.StatusCreated, u)
	}
}

// Обработчик создания сессии
func (s *server) handleSessionsCreate() http.HandlerFunc {
	// Структура "запрос"
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Парсинг запроса
		req := new(request)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		// Аутентификация пользователя
		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		// Создание сессии
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		// Созранение ID пользователя в сессии
		session.Values["user_id"] = u.ID
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
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
