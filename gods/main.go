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

	list := arraylist.New() // На основе массива
	// list := singlylinkedlist.New() // На основе связного списка
	// list := doublylinkedlist.New() // На основе двусвязного списка

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

	fmt.Println("Замена элементов местами")
	list.Swap(0, 1)
	list.Swap(0, 2)
	fmt.Println(list.Values())

	fmt.Println("Удаление элементов")
	list.Remove(2)
	list.Remove(1)
	list.Remove(1)
	fmt.Println(list.Values())

	fmt.Println("Вставка элементов в начало")
	list.Insert(0, "b")
	list.Insert(0, "a")
	fmt.Println(list.Values())

	fmt.Println("Очистка всего списка")
	fmt.Printf("размер: %d, пуст: %t\n", list.Size(), list.Empty())
	list.Clear()
	fmt.Printf("размер: %d, пуст: %t\n", list.Size(), list.Empty())
}
