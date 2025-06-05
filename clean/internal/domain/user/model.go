package user

import "errors"

// Структура "пользователь"
type User struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Address string `json:"address,omitempty"`
	Email   string `json:"email,omitempty"`
	Age     int    `json:"age"`
}

// Проверить почту
func (u *User) CheckEmail() error {
	if u.Email == "" {
		return errors.New("email не задан")
	}
	return nil
}

// Проверить возраст
func (u *User) CheckAge(minAge int) error {
	if u.Age < minAge {
		return errors.New("возраст не достигнут")
	}
	return nil
}
