package author

import (
	"context"
	"errors"

	"go-learn/clean/internal/domain/author"
	"go-learn/clean/pkg/client/memory"
)

// Структура "хранилище"
type storage struct {
	client *memory.Client
}

// Конструктор хранилища
func NewStorage() author.Storage {
	return &storage{
		client: memory.NewClient(),
	}
}

// Получение конкретного автора
func (s *storage) GetOne(ctx context.Context, uuid string) (*author.Author, error) {
	value, ok := s.client.Get(uuid)
	if !ok {
		return nil, errors.New("не найдено")
	}
	return value.(*author.Author), nil
}

// Получение всех авторов
func (s *storage) GetAll(ctx context.Context, limit, offset int) ([]*author.Author, error) {
	authors := make([]*author.Author, 0, limit)
	for _, value := range s.client.Values() {
		authors = append(authors, value.(*author.Author))
	}
	if offset > len(authors) {
		offset = len(authors)
	}
	authors = authors[offset:]
	if len(authors) > limit {
		authors = authors[:limit]
	}
	return authors, nil
}

// Создание автора
func (s *storage) Create(ctx context.Context, author *author.Author) (*author.Author, error) {
	if ok := s.client.Put(author.UUID, author); !ok {
		return nil, errors.New("не сохранено")
	}
	return author, nil
}

// Обновление автора
func (s *storage) Update(ctx context.Context, author *author.Author) (*author.Author, error) {
	if _, ok := s.client.Get(author.UUID); !ok {
		return nil, errors.New("не найдено")
	}
	if ok := s.client.Put(author.UUID, author); !ok {
		return nil, errors.New("не обновлено")
	}
	return author, nil
}

// Удаление автора
func (s *storage) Delete(ctx context.Context, author *author.Author) error {
	if ok := s.client.Delete(author.UUID); !ok {
		return errors.New("не удалено")
	}
	return nil
}
