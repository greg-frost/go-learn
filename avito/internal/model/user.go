package model

// Структура "пользователь"
type User struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email"`
	Role  Role   `json:"role"`
}

// Структура "объект передачи информации о пользователе"
type UserDTO struct {
	Role Role `json:"role"`
}

// Роль пользователя
type Role string

// Допустимые значения роли
const (
	RoleEmployee  Role = "employee"
	RoleModerator Role = "moderator"
)

// Проверка валидности роли
func (r Role) Valid() bool {
	switch r {
	case RoleEmployee, RoleModerator:
		return true
	default:
		return false
	}
}
