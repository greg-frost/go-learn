package book

// Структура "пользователь"
type User struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Address string `json:"address,omitempty"`
	Email   string `json:"email,omitempty"`
	Age     int    `json:"age"`
}
