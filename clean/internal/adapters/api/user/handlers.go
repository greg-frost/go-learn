package user

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"go-learn/clean/internal/adapters/api"
	"go-learn/clean/internal/domain/user"

	"github.com/julienschmidt/httprouter"
)

const (
	usersURL = "/users"          // URL конкретного пользователя
	userURL  = "/users/:user_id" // URL списка пользователей
)

// Структура "обработчик"
type handler struct {
	service user.Service
}

// Конструктор обработчика
func NewHandler(service user.Service) api.Handler {
	return &handler{
		service: service,
	}
}

// Регистрация маршрутов
func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetAllUsers)
	router.GET(userURL, h.GetUserByUUID)
}

// Получение всех пользователей
func (h *handler) GetAllUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	users, _ := h.service.GetAllUsers(context.Background(), limit, offset)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// Получение конкретного пользователя
func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	uuid := params.ByName("user_id")
	user, _ := h.service.GetUserByUUID(context.Background(), uuid)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
