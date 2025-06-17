package book

// Структура "создание книги"
type CreateBookDTO struct {
	Title      string `json:"title"`
	AuthorUUID string `json:"author_uuid"`
	Year       int    `json:"year"`
}

// Структура "обновление книги"
type UpdateBookDTO struct {
	UUID       string `json:"uuid"`
	Title      string `json:"title"`
	AuthorUUID string `json:"author_uuid"`
	Year       int    `json:"year"`
	Busy       bool   `json:"busy"`
	OwnerUUID  string `json:"owner_uuid"`
}
