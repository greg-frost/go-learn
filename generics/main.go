package main

import (
	"fmt"
)

// Структура список произвольных элементов
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

	/* Функции-дженерики */

	sis := 15
	si := []int{10, 20, 15, -10}
	fmt.Println("Поиск", sis, "в", si, ":", Index(si, sis))

	sss := "hello"
	ss := []string{"foo", "bar", "baz"}
	fmt.Println("Поиск", sss, "в", ss, ":", Index(ss, sss))

	/* Типы-дженерики */

	fmt.Println()
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
