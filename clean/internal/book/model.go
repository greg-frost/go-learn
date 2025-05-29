package book

// Структура "книга"
type Book struct {
	UUID   string `json:"uuid"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year,omitempty"`
	Busy   bool   `json:"busy"`
	User   string `json:"user,omitempty"`
}
