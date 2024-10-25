package main

import (
	"container/heap"
	"fmt"
)

/* Библиотечная */

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

/* Собственная */

// Минимальная
type MinHeap []int

func (h MinHeap) sink(i int) {
	n := len(h)
	k := i
	j := 2*k + 1

	for j < n {
		if j < n-1 && h[j+1] < h[j] {
			j++
		}
		if h[k] <= h[j] {
			return
		}

		h[k], h[j] = h[j], h[k]
		k = j
		j = 2*k + 1
	}
}

func (h MinHeap) Heapify() {
	n := len(h)
	for i := (n - 1) / 2; i >= 0; i-- {
		h.sink(i)
	}
}

func (h *MinHeap) Push(v int) {
	*h = append(*h, v)
	h.Heapify()
}

func (h *MinHeap) Pop() (int, bool) {
	if len(*h) == 0 {
		return 0, false
	}

	v := (*h)[0]
	*h = (*h)[1:]
	h.Heapify()

	return v, true
}

// Максимальная
type MaxHeap []int

func (h MaxHeap) sink(i int) {
	n := len(h)
	k := i
	j := 2*k + 1

	for j < n {
		if j < n-1 && h[j+1] > h[j] {
			j++
		}
		if h[k] >= h[j] {
			return
		}

		h[k], h[j] = h[j], h[k]
		k = j
		j = 2*k + 1
	}
}

func (h MaxHeap) Heapify() {
	n := len(h)
	for i := (n - 1) / 2; i >= 0; i-- {
		h.sink(i)
	}
}

func (h *MaxHeap) Push(v int) {
	*h = append(*h, v)
	h.Heapify()
}

func (h *MaxHeap) Pop() (int, bool) {
	if len(*h) == 0 {
		return 0, false
	}

	v := (*h)[0]
	*h = (*h)[1:]
	h.Heapify()

	return v, true
}

func main() {
	fmt.Println(" \n[ КУЧА ]\n ")

	/* Библиотечная куча */

	fmt.Println("Библиотечная")
	fmt.Println("------------")
	fmt.Println()

	h1 := &IntHeap{}
	heap.Init(h1)

	// Добавление

	fmt.Println("Push: 10")
	heap.Push(h1, 10)

	fmt.Println("Push: 5")
	heap.Push(h1, 5)

	fmt.Println("Push: 15")
	heap.Push(h1, 15)

	// Состояние
	fmt.Println()
	fmt.Println("Heap:", *h1)
	fmt.Println()

	// Извлечение

	v := heap.Pop(h1).(int)
	fmt.Println("Min:", v)

	v = heap.Pop(h1).(int)
	fmt.Println("Min:", v)

	v = heap.Pop(h1).(int)
	fmt.Println("Min:", v)

	// Состояние
	fmt.Println()
	fmt.Println("Heap:", *h1)
	fmt.Println()

	/* Собственная куча */

	fmt.Println("Собственная")
	fmt.Println("-----------")
	fmt.Println()

	h2 := MinHeap{}

	// Добавление

	fmt.Println("Push: 3")
	h2.Push(3)

	fmt.Println("Push: 5")
	h2.Push(5)

	fmt.Println("Push: 1")
	h2.Push(1)

	// Состояние
	fmt.Println()
	fmt.Println("Heap:", h2)
	fmt.Println()

	// Извлечение

	v, ok := h2.Pop()
	fmt.Println("Min:", v, ok)

	v, ok = h2.Pop()
	fmt.Println("Min:", v, ok)

	v, ok = h2.Pop()
	fmt.Println("Min:", v, ok)

	v, ok = h2.Pop()
	fmt.Println("Min:", v, ok)

	// Состояние
	fmt.Println()
	fmt.Println("Heap:", h2)
}
