package user

import "golang.org/x/crypto/bcrypt"

// Структура "пользователь"
type User struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Email        string `json:"email" bson:"email"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"`
}

// Шифрование пароля
func (u User) encryptPassword(s string) error {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(b)
	return nil
}

// Структура "создаваемый пользователь"
type CreateUserDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Структура "обновляемый пользователь"
type UpdateUserDTO struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
