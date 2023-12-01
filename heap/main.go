package main

import (
	"fmt"
)

type Heap []int

func (h *Heap) Push(v int) {
	pos := 0
	for pos < len(*h) && v > (*h)[pos] {
		pos++
	}

	*h = append(*h, v)

	if pos < len(*h)-1 {
		copy((*h)[pos+1:], (*h)[pos:])
		(*h)[pos] = v
	}
}

func (h *Heap) PopMin() (int, bool) {
	if len(*h) == 0 {
		return 0, false
	}

	v := (*h)[0]
	*h = (*h)[1:]

	return v, true
}

func (h *Heap) PopMax() (int, bool) {
	if len(*h) == 0 {
		return 0, false
	}

	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return v, true
}

func main() {
	fmt.Println(" \n[ КУЧА ]\n ")

	heap := Heap{}

	// Добавление

	fmt.Println("Push: 3")
	heap.Push(3)

	fmt.Println("Push: 5")
	heap.Push(5)

	fmt.Println("Push: 1")
	heap.Push(1)

	// Состояние
	fmt.Println()
	fmt.Println("Heap:", heap)
	fmt.Println()

	// Извлечение

	v, ok := heap.PopMin()
	fmt.Println("Min:", v, ok)

	v, ok = heap.PopMax()
	fmt.Println("Max:", v, ok)

	v, ok = heap.PopMin()
	fmt.Println("Min:", v, ok)

	v, ok = heap.PopMax()
	fmt.Println("Max:", v, ok)

	// Состояние
	fmt.Println()
	fmt.Println("Heap:", heap)
}
