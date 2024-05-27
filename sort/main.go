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

// Функция длины для сортировки
func (this ByName) Len() int {
	return len(this)
}

// Функция поиска меньшего для сортировки
func (this ByName) Less(i, j int) bool {
	return this[i].firstName < this[j].firstName
}

// Функция замены элементов для сортировки
func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

// Сортировка по возрасту
type ByAge []Person

// Функция длины для сортировки
func (this ByAge) Len() int {
	return len(this)
}

// Функция поиска меньшего для сортировки
func (this ByAge) Less(i, j int) bool {
	return this[i].age < this[j].age
}

// Функция замены элементов для сортировки
func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	fmt.Println(" \n[ СОРТИРОВКА ]\n ")

	/* Простая сортировка */

	fmt.Println("Простая сортировка")
	fmt.Println("------------------")
	fmt.Println()

	nums := []int{3, 1, 7, 4, 2, 6, 5, 10, 8, 9}

	fmt.Println("До сортировки:")
	fmt.Println(nums)
	fmt.Println()

	sort.Ints(nums)

	fmt.Println("После сортировки:")
	fmt.Println(nums)

	/* Сложная сортировка */

	fmt.Println()
	fmt.Println("Сложная сортировка")
	fmt.Println("------------------")
	fmt.Println()

	var people = []Person{
		{"Charles", "Bukowski", 27},
		{"Ada", "Wong", 21},
		{"Bob", "Marley", 18},
	}

	fmt.Println("До сортировки:")
	fmt.Println(people)
	fmt.Println()

	/* По имени */

	fmt.Println("Сортировка по имени:")
	sort.Sort(ByName(people))
	fmt.Println(people)
	fmt.Println()

	/* По фамилии */

	fmt.Println("Сортировка по фамилии:")
	sort.Slice(people, func(i, j int) bool {
		return people[i].lastName < people[j].lastName
	})
	fmt.Println(people)
	fmt.Println()

	/* По возрасту */

	fmt.Println("Сортировка по возрасту:")
	sort.Sort(ByAge(people))
	fmt.Println(people)
}
