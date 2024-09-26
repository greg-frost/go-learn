package main

// Структура "очередь"
type MyIntQueue struct {
	q []MyInt
}

// Конструктор очереди
func NewMyIntQueue() *MyIntQueue {
	return &MyIntQueue{
		q: []MyInt{},
	}
}

// Вставка значения в конец очереди
func (o *MyIntQueue) Insert(v MyInt) {
	o.q = append(o.q, v)
}

// Получение значения из начала очереди
func (o *MyIntQueue) Remove() MyInt {
	if len(o.q) == 0 {
		panic("Пусто!")
	}
	first := o.q[0]
	o.q = o.q[1:]
	return first
}