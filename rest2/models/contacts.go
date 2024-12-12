package models

import (
	"github.com/jinzhu/gorm"
)

// Структура "контакт"
type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserID uint   `json:"user_id"`
}
