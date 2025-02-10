package main

import (
	"fmt"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/utils"
)

func main() {
	fmt.Println(" \n[ GO DATA STRUCTURES (GODS) ]\n ")

	var value interface{}
	var ok bool

	fmt.Println("Список:")
	fmt.Println()

	list := arraylist.New() // Массив

	fmt.Println("Добавление элементов")
	list.Add("a")
	list.Add("c", "b")
	fmt.Println(list.Values())

	fmt.Println("Сортировка (как строк)")
	list.Sort(utils.StringComparator)
	fmt.Println(list.Values())

	fmt.Println("Получение по индексу")
	value, ok = list.Get(0)
	fmt.Println("0:", value, ok)
	value, ok = list.Get(3)
	fmt.Println("3:", value, ok)

	fmt.Println("Наличие элементов (всех)")
	ok = list.Contains("a", "b", "c")
	fmt.Println("a, b, c:", ok)
	ok = list.Contains("a", "b", "c", "d")
	fmt.Println("a, b, c, d:", ok)
}
