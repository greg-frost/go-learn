package model

import "time"

// Структура "товар"
type Product struct {
	ID          string    `json:"id,omitempty"`
	DateTime    time.Time `json:"dateTime,omitempty"`
	Type        Type      `json:"type"`
	ReceptionID string    `json:"receptionId"`
}

// Структура "объект передачи информации о товаре"
type ProductDTO struct {
	Type  Type   `json:"type"`
	PvzID string `json:"pvzId"`
}

// Тип товара
type Type string

// Допустимые значения типа
const (
	TypeElectronics Type = "электроника"
	TypeClothes     Type = "одежда"
	TypeShoes       Type = "обувь"
)

// Проверка валидности типа
func (t Type) Valid() bool {
	switch t {
	case TypeElectronics, TypeClothes, TypeShoes:
		return true
	default:
		return false
	}
}
