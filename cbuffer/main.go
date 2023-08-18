package main

import (
	"fmt"
)

// Структура "кольцевой буфер"
type CircularBuffer struct {
	values  []float64
	headIdx int
	tailIdx int
}

// Конструктор буфера
func NewCircularBuffer(size int) CircularBuffer {
	return CircularBuffer{values: make([]float64, size+1)}
}

// Добавление нового значения в буфер
func (b *CircularBuffer) AddValue(v float64) {
	b.values[b.tailIdx] = v
	b.tailIdx = (b.tailIdx + 1) % cap(b.values)
	if b.tailIdx == b.headIdx {
		b.headIdx = (b.headIdx + 1) % cap(b.values)
	}
}

// Длина буфера
func (b CircularBuffer) GetSize() int {
	if b.tailIdx < b.headIdx {
		return b.tailIdx + cap(b.values) - b.headIdx
	}

	return b.tailIdx - b.headIdx
}

// Значения буфера с сохранением порядка записи
func (b CircularBuffer) GetValues() (retValues []float64) {
	for i := b.headIdx; i != b.tailIdx; i = (i + 1) % cap(b.values) {
		retValues = append(retValues, b.values[i])
	}

	return
}

// Установка значения по индексу (не надо так делать!)
func (b CircularBuffer) SetValueByIdx(idx int, v float64) {
	b.values[idx] = v
}

// Обработчик функции
func Handle(v float64, f func(float64)) {
	f(v)
}

// Структура "расширенный кольцевой буфер"
type ExtendedCircularBuffer struct {
	CircularBuffer
}

// Добавление нескольких значений
func (cb *ExtendedCircularBuffer) AddValues(vals ...float64) {
	for _, val := range vals {
		cb.AddValue(val)
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

	/* Простой буфер */

	buf := NewCircularBuffer(4)
	for i := 0; i < 7; i++ {
		if i > 0 {
			buf.AddValue(float64(i))
		}
		fmt.Printf("[%d]: %v\n", buf.GetSize(), buf.GetValues())
	}

	/* Изменение значений */

	buf.SetValueByIdx(0, -1.0)
	buf.SetValueByIdx(1, -2.0)
	Handle(10.0, buf.AddValue)
	Handle(20.0, buf.AddValue)
	Handle(30.0, buf.AddValue)

	fmt.Println()
	fmt.Println("Новые значения:")
	fmt.Println(buf.values)

	/* Расширенный буфер */

	fmt.Println()
	fmt.Println("Расширенный буфер:")

	extBuf := NewExtendedCircularBuffer(5)
	extBuf.AddValues(1, 2, 3, 4, 5)
	fmt.Printf("[%d]: %v\n", extBuf.GetSize(), extBuf.GetValues())
}
