package store

// Интерфейс "хранилище"
type Store interface {
	User() UserRepository
}
