package book

import (
	"go-learn/clean/internal/domain/author"
)

// Структура "создание книги"
type CreateBookDTO struct {
	Title  string        `json:"title"`
	Author author.Author `json:"author"`
	Year   int           `json:"year"`
}

// Структура "обновление книги"
type UpdateBookDTO struct {
	UUID    string        `json:"uuid"`
	Title   string        `json:"title"`
	Author  author.Author `json:"author"`
	Year    int           `json:"year"`
	Busy    bool          `json:"busy"`
	OwnerID string        `json:"owner_id"`
}

// Структура "удаление книги"
type DeleteBookDTO struct {
	UUID string `json:"uuid"`
}
