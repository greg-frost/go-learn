package model

import "time"

// Структура "приемка"
type Reception struct {
	ID       string    `json:"id,omitempty"`
	DateTime time.Time `json:"dateTime"`
	PvzID    string    `json:"pvzId"`
	Status   Status    `json:"status"`
}

// Структура "объект передачи информации о приемке"
type ReceptionDTO struct {
	PvzID string `json:"pvzId"`
}

// Структура "результат выдачи приемки"
type ReceptionResult struct {
	Reception Reception `json:"reception"`
	Products  []Product `json:"products"`
}

// Статус приемки
type Status string

// Допустимые значения статуса
const (
	StatusInProgress Status = "in_progress"
	StatusClose      Status = "close"
)

// Проверка валидности статуса
func (s Status) Valid() bool {
	switch s {
	case StatusInProgress, StatusClose:
		return true
	default:
		return false
	}
}
