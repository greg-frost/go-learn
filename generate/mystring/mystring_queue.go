package main

// Структура "очередь"
type MyStringQueue struct {
	q []MyString
}

// Конструктор очереди
func NewMyStringQueue() *MyStringQueue {
	return &MyStringQueue{
		q: []MyString{},
	}
}

// Вставка значения в конец очереди
func (o *MyStringQueue) Insert(v MyString) {
	o.q = append(o.q, v)
}

// Получение значения из начала очереди
func (o *MyStringQueue) Remove() MyString {
	if len(o.q) == 0 {
		panic("Пусто!")
	}
	first := o.q[0]
	o.q = o.q[1:]
	return first
}