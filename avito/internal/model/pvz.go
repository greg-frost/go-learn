package model

import "time"

// Структура "ПВЗ" (пункт выдачи заказов)
type PVZ struct {
	ID               string    `json:"id,omitempty"`
	RegistrationDate time.Time `json:"registrationDate,omitempty"`
	City             City      `json:"city"`
}

// Структура "объект передачи информации о ПВЗ"
type PvzDTO struct {
	City City `json:"city"`
}

// Структура "результат выдачи ПВЗ"
type PVZResult struct {
	PVZ        PVZ               `json:"pvz"`
	Receptions []ReceptionResult `json:"receptions"`
}

// Город ПВЗ
type City string

// Допустимые значения города
const (
	CityMoscow          City = "Москва"
	CitySaintPetersburg City = "Санкт-Петербург"
	CityKazan           City = "Казань"
)

// Проверка валидности города
func (c City) Valid() bool {
	switch c {
	case CityMoscow, CitySaintPetersburg, CityKazan:
		return true
	default:
		return false
	}
}
