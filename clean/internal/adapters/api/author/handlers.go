package author

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"go-learn/clean/internal/adapters/api"
	"go-learn/clean/internal/domain/author"

	"github.com/julienschmidt/httprouter"
)

const (
	authorsURL = "/authors"            // URL списка авторов
	authorURL  = "/authors/:author_id" // URL конкретного автора
)

// Структура "обработчик"
type handler struct {
	service author.Service
}

// Конструктор обработчика
func NewHandler(service author.Service) api.Handler {
	return &handler{
		service: service,
	}
}

// Регистрация маршрутов
func (h *handler) Register(router *httprouter.Router) {
	router.GET(authorURL, h.GetAuthorByUUID)
	router.GET(authorsURL, h.GetAllAuthors)
}

// Получение конкретного автора
func (h *handler) GetAuthorByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	uuid := params.ByName("author_id")
	author, _ := h.service.GetAuthorByUUID(context.Background(), uuid)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(author)
}

// Получение всех авторов
func (h *handler) GetAllAuthors(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	authors, _ := h.service.GetAllAuthors(context.Background(), limit, offset)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authors)
}
