package book

import "errors"

// Структура "книга"
type Book struct {
	UUID   string `json:"uuid"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year,omitempty"`
	Busy   bool   `json:"busy"`
	User   string `json:"user,omitempty"`
}

// Взять книгу
func (b *Book) Take(user string) error {
	if b.Busy {
		return errors.New("книга занята")
	}
	b.User = user
	b.Busy = true
	return nil
}
