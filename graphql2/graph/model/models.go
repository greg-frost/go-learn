package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// Тип "номер"
type Num int

// Тип "временная метка"
type Timestamp time.Time

// Структура "видео"
type Video struct {
	ID          Num       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserID      int       `json:"-"`
	URL         string    `json:"url"`
	CreatedAt   Timestamp `json:"createdAt"`
	Genre       *Genre    `json:"genre,omitempty"`
}

// Структура "новое видео"
type NewVideo struct {
	ID          *Num   `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      int    `json:"userId"`
	URL         string `json:"url"`
	Genre       *Genre `json:"genre,omitempty"`
}

// Маршаллинг номера
func MarshalNum(id Num) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(fmt.Sprintf("%d", id)))
	})
}

// Демаршаллинг номера
func UnmarshalNum(v interface{}) (Num, error) {
	id, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("номер должен быть строкой")
	}
	i, e := strconv.Atoi(id)
	return Num(i), e
}

// Маршаллинг временной метки
func MarshalTimestamp(t Timestamp) graphql.Marshaler {
	timestamp := time.Time(t).Unix() * 1000

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(timestamp, 10))
	})
}

// Демаршаллинг временной метки
func UnmarshalTimestamp(v interface{}) (Timestamp, error) {
	if s, ok := v.(int); ok {
		return Timestamp(time.Unix(int64(s), 0)), nil
	}
	return Timestamp{}, fmt.Errorf("не удалось преобразовать дату")
}
