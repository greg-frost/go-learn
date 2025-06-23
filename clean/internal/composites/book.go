package composites

import (
	handler "go-learn/clean/internal/adapters/api"
	api "go-learn/clean/internal/adapters/api/book"
	storage "go-learn/clean/internal/adapters/db/book"
	service "go-learn/clean/internal/domain/author"
	domain "go-learn/clean/internal/domain/book"
)

// Структура "композит книги"
type BookComposite struct {
	Storage domain.Storage
	Service api.Service
	Handler handler.Handler
}

// Конструктор композита
func NewBookComposite(authorService service.Service) (*BookComposite, error) {
	storage := storage.NewStorage()
	service := domain.NewService(storage, authorService)
	handler := api.NewHandler(service)

	return &BookComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
