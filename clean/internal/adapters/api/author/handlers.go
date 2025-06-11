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
	authorsURL = "/authors"            // URL конкретного автора
	authorURL  = "/authors/:author_id" // URL списка авторов
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
	router.GET(authorsURL, h.GetAllAuthors)
	router.GET(authorURL, h.GetAuthorByUUID)
}

// Получение всех авторов
func (h *handler) GetAllAuthors(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	authors, _ := h.service.GetAllAuthors(context.Background(), limit, offset)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authors)
}

// Получение конкретного автора
func (h *handler) GetAuthorByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	uuid := params.ByName("author_id")
	author, _ := h.service.GetAuthorByUUID(context.Background(), uuid)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(author)
}
