package main

import (
	"fmt"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/stacks/arraystack"
	"github.com/emirpasic/gods/utils"
)

func main() {
	fmt.Println(" \n[ GO DATA STRUCTURES (GODS) ]\n ")

	var value interface{}
	var ok bool

	/* Список (List) */

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

	fmt.Println("Полная очистка")
	fmt.Printf("размер: %d, пуст: %t\n", list.Size(), list.Empty())
	list.Clear()
	fmt.Printf("размер: %d, пуст: %t\n", list.Size(), list.Empty())
	fmt.Println()

	/* Множество (Set) */

	fmt.Println("Множество:")
	fmt.Println()

	set := hashset.New() // На основе хеш-таблицы (случайный порядок)
	// set := treeset.NewWithIntComparator() // На основе дерева (упорядочено)
	// set := linkedhashset.New() // На основе хеш-таблицы и списка (в порядке вставки)

	fmt.Println("Добавление элементов")
	set.Add(1)
	set.Add(2, 2, 5, 4, 3)
	fmt.Println(set.Values())

	fmt.Println("Наличие элементов (всех)")
	ok = set.Contains(1)
	fmt.Println("1:", ok)
	ok = set.Contains(2, 5)
	fmt.Println("2, 5:", ok)
	ok = set.Contains(3, 6)
	fmt.Println("3, 6:", ok)

	fmt.Println("Удаление элементов")
	set.Remove(4)
	set.Remove(2, 3, 4)
	fmt.Println(set.Values())

	fmt.Println("Полная очистка")
	fmt.Printf("размер: %d, пуст: %t\n", set.Size(), set.Empty())
	set.Clear()
	fmt.Printf("размер: %d, пуст: %t\n", set.Size(), set.Empty())
	fmt.Println()

	/* Стек (Stack) */

	fmt.Println("Стек:")
	fmt.Println()

	stack := arraystack.New() // На основе массива
	// stack := linkedliststack.New() // На основе связного списка

	fmt.Println("Добавление элементов")
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Values())

	fmt.Println("Просмотр и извлечение")
	value, ok = stack.Peek()
	fmt.Println("Peek:", value, ok)
	value, ok = stack.Pop()
	fmt.Println("Pop:", value, ok)
	value, ok = stack.Pop()
	fmt.Println("Pop:", value, ok)
	value, ok = stack.Pop()
	fmt.Println("Pop:", value, ok)

	fmt.Println("Полная очистка")
	stack.Push(3)
	fmt.Printf("размер: %d, пуст: %t\n", stack.Size(), stack.Empty())
	stack.Clear()
	fmt.Printf("размер: %d, пуст: %t\n", stack.Size(), stack.Empty())
	fmt.Println()

	/* Карта (Map) */

	fmt.Println("Карта:")
	fmt.Println()

	m := hashmap.New() // На основе хеш-таблицы (случайный порядок)

	fmt.Println("Добавление элементов")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a")
	m.Put(3, "c")
	b, _ := m.ToJSON()
	fmt.Println(string(b))

	fmt.Println("Получение по ключу")
	value, ok = m.Get(2)
	fmt.Println("2:", value, ok)
	value, ok = m.Get(4)
	fmt.Println("4:", value, ok)

	fmt.Println("Ключи и значения")
	fmt.Println(m.Keys(), m.Values())
	fmt.Println("Удаление элементов")
	m.Remove(3)
	fmt.Println(m.Values())

	fmt.Println("Полная очистка")
	fmt.Printf("размер: %d, пуст: %t\n", m.Size(), m.Empty())
	m.Clear()
	fmt.Printf("размер: %d, пуст: %t\n", m.Size(), m.Empty())
}
