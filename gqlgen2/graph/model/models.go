package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type Num int
type Timestamp time.Time

type Video struct {
	ID          Num           `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	User        *User         `json:"user"`
	URL         string        `json:"url"`
	CreatedAt   Timestamp     `json:"createdAt"`
	Screenshots []*Screenshot `json:"screenshots,omitempty"`
	Related     []*Video      `json:"related"`
}

func MarshalNum(id Num) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(fmt.Sprintf("%d", id)))
	})
}

func UnmarshalNum(v interface{}) (Num, error) {
	id, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("ID должен быть строкой")
	}
	i, e := strconv.Atoi(id)
	return Num(i), e
}

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
