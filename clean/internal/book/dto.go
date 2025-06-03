package book

// Структура "создание книги"
type CreateBookDTO struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

// Структура "обновление книги"
type UpdateBookDTO struct {
	UUID   string `json:"uuid"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	Busy   bool   `json:"busy"`
	User   string `json:"user"`
}
