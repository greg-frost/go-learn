package book

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-learn/clean/internal/adapters/api"
	"go-learn/clean/internal/domain/book"

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
	router.POST(booksURL, h.CreateBook)
	router.PUT(bookURL, h.UpdateBook)
	router.DELETE(bookURL, h.DeleteBook)
}

// Получение конкретной книги
func (h *handler) GetBookByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	uuid := params.ByName("book_id")
	book, err := h.service.GetBookByUUID(r.Context(), uuid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "ошибка поиска книги: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

const limitDefault = 10 // Ограничение выборки по умолчанию

// Получение всех книг
func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 0 {
		limit = limitDefault
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	books, err := h.service.GetAllBooks(r.Context(), limit, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "ошибка поиска книг: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

// Создание книги
func (h *handler) CreateBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var bookDTO book.CreateBookDTO
	if err := json.NewDecoder(r.Body).Decode(&bookDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ошибка парсинга данных: %v", err)
		return
	}

	book, err := h.service.CreateBook(r.Context(), &bookDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "ошибка создания книги: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// Обновление книги
func (h *handler) UpdateBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var bookDTO book.UpdateBookDTO
	if err := json.NewDecoder(r.Body).Decode(&bookDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ошибка парсинга данных: %v", err)
		return
	}
	bookDTO.UUID = params.ByName("book_id")

	book, err := h.service.UpdateBook(r.Context(), &bookDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "ошибка обновления книги: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)

}

// Удаление книги
func (h *handler) DeleteBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	uuid := params.ByName("book_id")

	err := h.service.DeleteBook(r.Context(), uuid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "ошибка удаления книги: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
