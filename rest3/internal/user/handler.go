package user

import (
	"encoding/json"
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
	service Service
	logger  *logger.Logger
}

// Конструктор обработчика
func NewHandler(service Service, logger *logger.Logger) handlers.Handler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

// Регистрация обработчиков
func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, apperror.Middleware(h.GetList))
	router.GET(userURL, apperror.Middleware(h.GetUserByID))
	router.POST(usersURL, apperror.Middleware(h.CreateUser))
	router.PUT(userURL, apperror.Middleware(h.UpdateUser))
	//router.PATCH(userURL, apperror.Middleware(h.EditUser))
	router.DELETE(userURL, apperror.Middleware(h.DeleteUser))
}

// Получение списка пользователей
func (h *handler) GetList(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	h.logger.Info("Список пользователей")
	users, err := h.service.GetUsers(r.Context())
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
	return nil
}

// Получение пользователя по ID
func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	h.logger.Info("Получение пользователя по ID")
	uuid := p.ByName("uuid")
	user, err := h.service.GetUserByID(r.Context(), uuid)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
	return nil
}

// Создание пользователя
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	w.WriteHeader(http.StatusCreated)
	h.logger.Info("Создание пользователя")
	w.Write([]byte("Создание пользователя"))
	return nil
	// return fmt.Errorf("API error")
}

// Полное обновление пользователя
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	w.WriteHeader(http.StatusNoContent)
	h.logger.Info("Полное обновление пользователя")
	w.Write([]byte("Полное обновление пользователя"))
	return nil
	// return apperror.NewAppError(nil, "внутренняя ошибка API", "user is immutable", "US-000004")
}

// Частичное обновление пользователя
// func (h *handler) EditUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
// 	w.WriteHeader(http.StatusNoContent)
// 	h.logger.Info("Частичное обновление пользователя")
// 	w.Write([]byte("Частичное обновление пользователя"))
// 	return nil
// }

// Удаление пользователя
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	w.WriteHeader(http.StatusNoContent)
	h.logger.Info("Удаление пользователя")
	w.Write([]byte("Удаление пользователя"))
	return nil
}
