package main

import (
	"container/heap"
	"fmt"
)

/* Моя */

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

/* Стандартная */

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	fmt.Println(" \n[ КУЧА ]\n ")

	/* Моя */

	fmt.Println("Моя")
	fmt.Println("---")
	fmt.Println()

	h1 := Heap{}

	// Добавление

	fmt.Println("Push: 3")
	h1.Push(3)

	fmt.Println("Push: 5")
	h1.Push(5)

	fmt.Println("Push: 1")
	h1.Push(1)

	// Состояние
	fmt.Println()
	fmt.Println("Heap:", h1)
	fmt.Println()

	// Извлечение

	v, ok := h1.PopMin()
	fmt.Println("Min:", v, ok)

	v, ok = h1.PopMax()
	fmt.Println("Max:", v, ok)

	v, ok = h1.PopMin()
	fmt.Println("Min:", v, ok)

	v, ok = h1.PopMax()
	fmt.Println("Max:", v, ok)

	// Состояние
	fmt.Println()
	fmt.Println("Heap:", h1)
	fmt.Println()

	/* Стандартная */

	fmt.Println("Стандартная")
	fmt.Println("-----------")
	fmt.Println()

	h2 := &IntHeap{}
	heap.Init(h2)

	// Добавление

	fmt.Println("Push: 10")
	heap.Push(h2, 10)

	fmt.Println("Push: 5")
	heap.Push(h2, 5)

	fmt.Println("Push: 15")
	heap.Push(h2, 15)

	// Состояние
	fmt.Println()
	fmt.Println("Heap:", *h2)
	fmt.Println()

	// Извлечение

	v = heap.Pop(h2).(int)
	fmt.Println("Min:", v)

	v = heap.Pop(h2).(int)
	fmt.Println("Min:", v)

	v = heap.Pop(h2).(int)
	fmt.Println("Min:", v)

	// Состояние
	fmt.Println()
	fmt.Println("Heap:", *h2)

}
