package apperror

import "encoding/json"

// Ошибки
var (
	ErrNotAuth  = NewAppError(nil, "не авторизовано", "user not authorized", "US-000002")
	ErrNotFound = NewAppError(nil, "не найдено", "entity not found", "US-000003")
)

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

// Маршаллинг
func (e *AppError) Marshal() []byte {
	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return b
}

// Системная ошибка
func systemError(err error) *AppError {
	return NewAppError(err, "внутренняя системная ошибка", err.Error(), "US-000001")
}
