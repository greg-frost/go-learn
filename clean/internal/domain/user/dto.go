package user

// Структура "создание пользователя"
type CreateUserDTO struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
}

// Структура "обновление пользователя"
type UpdateUserDTO struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
}
