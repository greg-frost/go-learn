package author

// Структура "создание автора"
type CreateAuthorDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Структура "обновление автора"
type UpdateAuthorDTO struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Структура "удаление автора"
type DeleteAuthorDTO struct {
	UUID string `json:"uuid"`
}
