package models

import (
	u "go-learn/rest2/utils"
	"log"

	"github.com/jinzhu/gorm"
)

// Структура "контакт"
type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserID uint   `json:"user_id"`
}

// Валидация контакта
func (c *Contact) Validate() (map[string]interface{}, bool) {
	if c.Name == "" {
		return u.Message(false, "Имя контакта необходимо"), false
	}

	if c.Phone == "" {
		return u.Message(false, "Телефон контакта необходим"), false
	}

	if c.UserID < 0 {
		return u.Message(false, "Пользователь не распознан"), false
	}

	return u.Message(true, "Валидация контакта пройдена"), true
}

// Создание контакта
func (c *Contact) Create() map[string]interface{} {
	if resp, ok := c.Validate(); !ok {
		return resp
	}

	DB().Create(c)

	resp := u.Message(true, "Контакт создан")
	resp["contact"] = c
	return resp
}

// Получение контакта
func GetContact(id uint) *Contact {
	contact := &Contact{}

	err := DB().Table("contacts").Where("id=?", id).First(contact).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return contact
}

// Получение контактов
func GetContacts(user uint) []*Contact {
	contacts := make([]*Contact, 0)

	err := DB().Table("contacts").Where("user_id=?", user).Find(&contacts).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return contacts
}
