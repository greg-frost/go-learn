package user

import (
	"fmt"
	"net/http"

	"go-learn/rest3/internal/apperror"
	"go-learn/rest3/internal/handlers"
	"go-learn/rest3/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

const (
	usersURL = "/users"       // URL списка пользователей
	userURL  = "/users/:uuid" // URL конкретного пользователя
)

// Структура "обработчик"
type handler struct {
	logger *logger.Logger
}

// Конструктор обработчика
func NewHandler(logger *logger.Logger) handlers.Handler {
	return &handler{logger: logger}
}

// Регистрация обработчиков
func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByID))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodPut, userURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))
}

// Получение списка пользователей
func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	// w.WriteHeader(200)
	// h.logger.Info("Список пользователей")
	// w.Write([]byte("Список пользователей"))
	// return nil
	return apperror.ErrNotFound
}

// Получение пользователя по ID
func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) error {
	// w.WriteHeader(200)
	// h.logger.Info("Получение пользователя по ID")
	// w.Write([]byte("Получение пользователя по ID"))
	// return nil
	return apperror.ErrNotAuth
}

// Создание пользователя
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	// w.WriteHeader(201)
	// h.logger.Info("Создание пользователя")
	// w.Write([]byte("Создание пользователя"))
	// return nil
	return fmt.Errorf("API error")
}

// Полное обновление пользователя
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	// w.WriteHeader(204)
	// h.logger.Info("Полное обновление пользователя")
	// w.Write([]byte("Полное обновление пользователя"))
	// return nil
	return apperror.NewAppError(nil, "внутренняя ошибка API", "user is immutable", "US-000004")
}

// Частичное обновление пользователя
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	h.logger.Info("Частичное обновление пользователя")
	w.Write([]byte("Частичное обновление пользователя"))
	return nil
}

// Удаление пользователя
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	h.logger.Info("Удаление пользователя")
	w.Write([]byte("Удаление пользователя"))
	return nil
}
