package composites

import (
	handler "go-learn/clean/internal/adapters/api"
	api "go-learn/clean/internal/adapters/api/book"
	storage "go-learn/clean/internal/adapters/db/book"
	domain "go-learn/clean/internal/domain/book"
)

// Структура "композит книги"
type BookComposite struct {
	Storage domain.Storage
	Service api.Service
	Handler handler.Handler
}

// Конструктор композита
func NewBookComposite(author *AuthorComposite) (*BookComposite, error) {
	storage := storage.NewStorage()
	service := domain.NewService(storage, author.Service)
	handler := api.NewHandler(service)

	return &BookComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
