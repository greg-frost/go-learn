package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// type Id int
type Timestamp time.Time

type Video struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	User        *User         `json:"user"`
	URL         string        `json:"url"`
	CreatedAt   Timestamp     `json:"createdAt"`
	Screenshots []*Screenshot `json:"screenshots,omitempty"`
	Related     []*Video      `json:"related"`
}

type NewVideo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      int    `json:"userId"`
	URL         string `json:"url"`
}

// func MarshalId(id Id) graphql.Marshaler {
// 	return graphql.WriterFunc(func(w io.Writer) {
// 		io.WriteString(w, strconv.Quote(fmt.Sprintf("%d", id)))
// 	})
// }

// func UnmarshalId(v interface{}) (Id, error) {
// 	id, ok := v.(string)
// 	if !ok {
// 		return 0, fmt.Errorf("ID должен быть строкой")
// 	}
// 	i, e := strconv.Atoi(id)
// 	return Id(i), e
// }

func MarshalTimestamp(t Timestamp) graphql.Marshaler {
	timestamp := time.Time(t).Unix() * 1000

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(timestamp, 10))
	})
}

func UnmarshalTimestamp(v interface{}) (Timestamp, error) {
	if s, ok := v.(int); ok {
		return Timestamp(time.Unix(int64(s), 0)), nil
	}
	return Timestamp{}, fmt.Errorf("не удалось преобразовать дату")
}
