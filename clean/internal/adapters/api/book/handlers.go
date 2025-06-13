package book

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"go-learn/clean/internal/adapters/api"

	"github.com/julienschmidt/httprouter"
)

const (
	booksURL = "/books"          // URL списка книг
	bookURL  = "/books/:book_id" // URL конкретной книги
)

// Структура "обработчик"
type handler struct {
	service Service
}

// Конструктор обработчика
func NewHandler(service Service) api.Handler {
	return &handler{
		service: service,
	}
}

// Регистрация маршрутов
func (h *handler) Register(router *httprouter.Router) {
	router.GET(bookURL, h.GetBookByUUID)
	router.GET(booksURL, h.GetAllBooks)
}

// Получение конкретной книги
func (h *handler) GetBookByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	uuid := params.ByName("book_id")
	book, _ := h.service.GetBookByUUID(context.Background(), uuid)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

// Получение всех книг
func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	books, _ := h.service.GetAllBooks(context.Background(), limit, offset)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
