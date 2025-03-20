package user

import (
	"net/http"

	"go-learn/rest3/internal/handlers"
	"go-learn/rest3/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

// Проверка соответствия интерфейсу
var _ handlers.Handler = new(handler)

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
	router.GET(usersURL, h.GetList)
	router.POST(usersURL, h.CreateUser)
	router.GET(userURL, h.GetUserByID)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

// Получение списка пользователей
func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	h.logger.Info("Список пользователей")
	w.Write([]byte("Список пользователей"))
}

// Создание пользователя
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	h.logger.Info("Создание пользователя")
	w.Write([]byte("Создание пользователя"))
}

// Получение пользователя по ID
func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	h.logger.Info("Получение пользователя по ID")
	w.Write([]byte("Получение пользователя по ID"))
}

// Полное обновление пользователя
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	h.logger.Info("Полное обновление пользователя")
	w.Write([]byte("Полное обновление пользователя"))
}

// Частичное обновление пользователя
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	h.logger.Info("Частичное обновление пользователя")
	w.Write([]byte("Частичное обновление пользователя"))
}

// Удаление пользователя
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	h.logger.Info("Удаление пользователя")
	w.Write([]byte("Удаление пользователя"))
}
