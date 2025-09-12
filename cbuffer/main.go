package main

import (
	"fmt"
)

// Структура "кольцевой буфер"
type CircularBuffer struct {
	values []float64
	head   int
	tail   int
}

// Конструктор буфера
func NewCircularBuffer(size int) CircularBuffer {
	return CircularBuffer{
		values: make([]float64, size+1),
	}
}

// Добавление нового значения в буфер
func (cb *CircularBuffer) Add(value float64) {
	cb.values[cb.tail] = value
	cb.tail = (cb.tail + 1) % cap(cb.values)
	if cb.tail == cb.head {
		cb.head = (cb.head + 1) % cap(cb.values)
	}
}

// Длина буфера
func (cb CircularBuffer) Size() int {
	if cb.tail < cb.head {
		return cb.tail + cap(cb.values) - cb.head
	}
	return cb.tail - cb.head
}

// Значения буфера с сохранением порядка записи
func (cb CircularBuffer) Values() []float64 {
	var res []float64
	for i := cb.head; i != cb.tail; i = (i + 1) % cap(cb.values) {
		res = append(res, cb.values[i])
	}
	return res
}

// Установка значения по индексу (не надо так делать!)
func (cb CircularBuffer) SetById(id int, value float64) {
	cb.values[id] = value
}

// Обработчик функции
func Handle(value float64, f func(float64)) {
	f(value)
}

// Структура "расширенный кольцевой буфер"
type ExtendedCircularBuffer struct {
	CircularBuffer
}

// Добавление нескольких значений
func (ecb *ExtendedCircularBuffer) Add(values ...float64) {
	for _, value := range values {
		ecb.CircularBuffer.Add(value)
	}
}

// Конструктор расширенного буфера
func NewExtendedCircularBuffer(size int) ExtendedCircularBuffer {
	return ExtendedCircularBuffer{
		CircularBuffer: NewCircularBuffer(size),
	}
}

func main() {
	fmt.Println(" \n[ КОЛЬЦЕВОЙ БУФЕР ]\n ")

	// Простой буфер
	buf := NewCircularBuffer(4)
	for i := 0; i < 7; i++ {
		if i > 0 {
			buf.Add(float64(i))
		}
		fmt.Printf("[%d]: %v\n", buf.Size(), buf.Values())
	}

	// Изменение значений
	buf.SetById(0, -1.0)
	buf.SetById(1, -2.0)
	Handle(10.0, buf.Add)
	Handle(20.0, buf.Add)
	Handle(30.0, buf.Add)
	fmt.Println()
	fmt.Println("Новые значения:")
	fmt.Println(buf.values)
	fmt.Println()

	// Расширенный буфер
	fmt.Println("Расширенный буфер:")
	extBuf := NewExtendedCircularBuffer(5)
	extBuf.Add(1, 2, 3, 4, 5)
	fmt.Printf("[%d]: %v\n", extBuf.Size(), extBuf.Values())
}
