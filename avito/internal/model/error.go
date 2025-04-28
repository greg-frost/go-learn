package model

import "fmt"

// Структура "ошибка"
type Error struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

// Вывод ошибки
func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
