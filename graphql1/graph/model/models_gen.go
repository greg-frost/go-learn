// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Query struct {
}

type Subscription struct {
}

type Time struct {
	UnixTime  int    `json:"unixTime"`
	TimeStamp string `json:"timeStamp"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
