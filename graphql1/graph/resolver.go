package graph

//go:generate go run github.com/99designs/gqlgen generate

import "golearn/graphql1/graph/model"

// Структура "обработчик"
type Resolver struct {
	todos []*model.Todo
}
