package author

import (
	"encoding/json"
	"fmt"
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
	router.POST(authorsURL, h.CreateAuthor)
	router.PUT(authorURL, h.UpdateAuthor)
	router.DELETE(authorURL, h.DeleteAuthor)
}

// Получение конкретного автора
func (h *handler) GetAuthorByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	uuid := params.ByName("author_id")
	author, err := h.service.GetAuthorByUUID(r.Context(), uuid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "ошибка поиска автора: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(author)
}

const (
	limitDefault  = 10 // Ограничение выборки по умолчанию
	offsetDefault = 0  // Смещение выборки по умолчанию
)

// Получение всех авторов
func (h *handler) GetAllAuthors(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 0 {
		limit = limitDefault
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || offset < 0 {
		offset = offsetDefault
	}

	authors, err := h.service.GetAllAuthors(r.Context(), limit, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "ошибка поиска авторов: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authors)
}

// Создание автора
func (h *handler) CreateAuthor(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var authorDTO author.CreateAuthorDTO
	if err := json.NewDecoder(r.Body).Decode(&authorDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ошибка парсинга данных: %v", err)
		return
	}

	author, err := h.service.CreateAuthor(r.Context(), &authorDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "ошибка создания автора: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}

// Обновление автора
func (h *handler) UpdateAuthor(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var authorDTO author.UpdateAuthorDTO
	if err := json.NewDecoder(r.Body).Decode(&authorDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ошибка парсинга данных: %v", err)
		return
	}
	authorDTO.UUID = params.ByName("author_id")

	author, err := h.service.UpdateAuthor(r.Context(), &authorDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "ошибка обновления автора: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)

}

// Удаление автора
func (h *handler) DeleteAuthor(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	uuid := params.ByName("author_id")

	err := h.service.DeleteAuthor(r.Context(), uuid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "ошибка удаления автора: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
