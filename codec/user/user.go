//go:generate codecgen -o user_generated.go user.go

package user

// Структура "пользователь"
type User struct {
	Name  string `codec:"User"`
	Email string `codec:",omitempty"`
}
