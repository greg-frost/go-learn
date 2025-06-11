package author

import "errors"

// Структура "автор"
type Author struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

// Проверить почту
func (u *Author) CheckEmail() error {
	if u.Email == "" {
		return errors.New("email не задан")
	}
	return nil
}
