package main

import (
	"fmt"
)

// Пользовательский тип
type MyType int

// Структура "очередь"
type MyTypeQueue struct {
	q []MyType
}

// Конструктор очереди
func NewMyTypeQueue() *MyTypeQueue {
	return &MyTypeQueue{
		q: []MyType{},
	}
}

// Вставка значения в конец очереди
func (o *MyTypeQueue) Insert(v MyType) {
	o.q = append(o.q, v)
}

// Получение значения из начала очереди
func (o *MyTypeQueue) Remove() MyType {
	if len(o.q) == 0 {
		panic("Пусто!")
	}
	first := o.q[0]
	o.q = o.q[1:]
	return first
}

func main() {
	fmt.Println(" \n[ КОДОГЕНЕРАЦИЯ ]\n ")

}
