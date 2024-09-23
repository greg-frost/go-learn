package graph

//go:generate go run github.com/99designs/gqlgen generate

import "go-learn/graphql2/graph/model"

// Структура "обработчик"
type Resolver struct {
	videos map[model.Num]model.Video
}

// Список подписчиков на событие публикации видео
var videoPublishedSubs = map[int]chan *model.Video{}
