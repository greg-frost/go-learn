package apperror

import "encoding/json"

// Структура "ошибка приложения"
type AppError struct {
	Err     error  `json:"-"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
	Code    string `json:"code"`
}

// Конструктор ошибки
func NewAppError(err error, message, reason, code string) *AppError {
	return &AppError{
		Err:     err,
		Message: message,
		Reason:  reason,
		Code:    code,
	}
}

// Стрингер
func (e *AppError) Error() string {
	return e.Message
}

// Разворачивание
func (e *AppError) Unwrap() error {
	return e.Err
}

// Маршалинг
func (e *AppError) Marshal() string {
	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return string(b)
}
