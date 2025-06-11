package book

import (
	"errors"

	"go-learn/clean/internal/domain/author"
)

// Структура "книга"
type Book struct {
	UUID    string        `json:"uuid"`
	Title   string        `json:"title"`
	Author  author.Author `json:"author"`
	Year    int           `json:"year,omitempty"`
	Busy    bool          `json:"busy"`
	OwnerID string        `json:"owner_id,omitempty"`
}

// Взять книгу
func (b *Book) Take(ownerID string) error {
	if b.Busy {
		return errors.New("книга занята")
	}
	b.OwnerID = ownerID
	b.Busy = true
	return nil
}
