package main

import (
	"fmt"
)

// Структура "список произвольных элементов"
type List[T any] struct {
	next *List[T]
	val  T
}

// Поиск элемента произвольного типа
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(" \n[ ДЖЕНЕРИКИ ]\n ")

	// Функции-дженерики
	fmt.Println("Функции:")
	seachInt := 15
	sliceInts := []int{10, 20, 15, -10}
	fmt.Printf("Поиск %d в %v: %d\n",
		seachInt, sliceInts, Index(sliceInts, seachInt))
	searchString := "hello"
	sliceStrings := []string{"foo", "bar", "baz"}
	fmt.Printf("Поиск %s в %v: %d\n",
		searchString, sliceStrings, Index(sliceStrings, searchString))
	fmt.Println()

	// Типы-дженерики
	fmt.Println("Связанные списки:")
	lastInt := new(List[int])
	lastInt.val = 1000
	firstInt := List[int]{lastInt, 0}
	fmt.Println(firstInt, lastInt)
	lastString := new(List[string])
	lastString.val = "World"
	firstString := List[string]{lastString, "Hello"}
	fmt.Println(firstString, lastString)
}
