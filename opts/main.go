package main

import (
	"fmt"
)

// Структура "предмет"
type Item struct {
	Default string
	Msg     string
	Code    int
}

// Тип "опция"
type option func(*Item)

// Конструктор предмета
func NewItem(opts ...option) *Item {
	i := &Item{
		Default: "req",
		Msg:     "OK",
		Code:    200,
	}

	for _, opt := range opts {
		opt(i)
	}

	return i
}

// Опция 1 (строковый параметр)
func Msg(val string) option {
	return func(i *Item) {
		i.Msg = val
	}
}

// Опция 2 (целый параметр)
func Code(val int) option {
	return func(i *Item) {
		i.Code = val
	}
}

func main() {
	fmt.Println(" \n[ ПАРАМЕТРЫ ]\n ")

	item1 := NewItem()
	item2 := NewItem(Code(201))
	item3 := NewItem(Msg("Bad"), Code(403))

	fmt.Printf("Item 1 = %v\n", *item1)
	fmt.Printf("Item 2 = %+v\n", *item2)
	fmt.Printf("Item 3 = %#v\n", *item3)
}
