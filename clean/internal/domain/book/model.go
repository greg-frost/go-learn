package book

import (
	"errors"

	"go-learn/clean/internal/domain/author"
)

// Структура "книга"
type Book struct {
	UUID      string        `json:"uuid"`
	Title     string        `json:"title"`
	Author    author.Author `json:"author"`
	Year      int           `json:"year,omitempty"`
	Busy      bool          `json:"busy"`
	OwnerUUID string        `json:"owner_uuid,omitempty"`
}

// Взять книгу
func (b *Book) Take(ownerUUID string) error {
	if b.Busy {
		return errors.New("книга занята")
	}
	b.OwnerUUID = ownerUUID
	b.Busy = true
	return nil
}
