package composites

import (
	handler "go-learn/clean/internal/adapters/api"
	api "go-learn/clean/internal/adapters/api/author"
	storage "go-learn/clean/internal/adapters/db/author"
	domain "go-learn/clean/internal/domain/author"
)

// Структура "композит автора"
type AuthorComposite struct {
	Storage domain.Storage
	Service domain.Service
	Handler handler.Handler
}

// Конструктор композита
func NewAuthorComposite() (*AuthorComposite, error) {
	storage := storage.NewStorage()
	service := domain.NewService(storage)
	handler := api.NewHandler(service)

	return &AuthorComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
