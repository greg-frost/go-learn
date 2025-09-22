package main

import (
	"fmt"
	"sort"
)

// Структура "человек"
type Person struct {
	firstName string
	lastName  string
	age       int
}

// Сортировка по имени
type ByName []Person

// Количество элементов
func (bn ByName) Len() int {
	return len(bn)
}

// Меньший элемент (по имени)
func (bn ByName) Less(i, j int) bool {
	return bn[i].firstName < bn[j].firstName
}

// Замена элементов
func (bn ByName) Swap(i, j int) {
	bn[i], bn[j] = bn[j], bn[i]
}

// Сортировка по возрасту
type ByAge []Person

// Количество элементов
func (ba ByAge) Len() int {
	return len(ba)
}

// Меньший элемент (по возрасту)
func (ba ByAge) Less(i, j int) bool {
	return ba[i].age < ba[j].age
}

// Замена элементов
func (ba ByAge) Swap(i, j int) {
	ba[i], ba[j] = ba[j], ba[i]
}

func main() {
	fmt.Println(" \n[ СОРТИРОВКА ]\n ")

	// Простая сортировка
	fmt.Println("Простая сортировка")
	fmt.Println("------------------")
	fmt.Println()

	fmt.Println("До сортировки:")
	nums := []int{3, 1, 7, 4, 2, 6, 5, 10, 8, 9}
	fmt.Println(nums)
	fmt.Println()

	fmt.Println("После сортировки:")
	sort.Ints(nums)
	fmt.Println(nums)
	fmt.Println()

	fmt.Println("В обратном порядке:")
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)
	fmt.Println()

	// Сложная сортировка
	fmt.Println("Сложная сортировка")
	fmt.Println("------------------")
	fmt.Println()

	fmt.Println("До сортировки:")
	var people = []Person{
		{"Charles", "Bukowski", 27},
		{"Ada", "Wong", 21},
		{"Bob", "Marley", 18},
	}
	fmt.Println(people)
	fmt.Println()

	// По имени
	fmt.Println("Сортировка по имени:")
	sort.Sort(ByName(people))
	fmt.Println(people)
	fmt.Println()

	// По фамилии
	fmt.Println("Сортировка по фамилии:")
	sort.Slice(people, func(i, j int) bool {
		return people[i].lastName < people[j].lastName
	})
	fmt.Println(people)
	fmt.Println()

	// По возрасту
	fmt.Println("Сортировка по возрасту:")
	sort.Sort(ByAge(people))
	fmt.Println(people)
}
