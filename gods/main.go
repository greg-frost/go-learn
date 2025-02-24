package main

import (
	"fmt"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/emirpasic/gods/queues/circularbuffer"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/stacks/arraystack"
	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
)

func main() {
	fmt.Println(" \n[ GO DATA STRUCTURES (GODS) ]\n ")

	var value interface{}
	var ok bool

	/* Список (List) */

	fmt.Println("Список")
	fmt.Println("------")
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

	fmt.Println("Множество")
	fmt.Println("---------")
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
	fmt.Printf("размер: %d, пусто: %t\n", set.Size(), set.Empty())
	set.Clear()
	fmt.Printf("размер: %d, пусто: %t\n", set.Size(), set.Empty())
	fmt.Println()

	/* Стек (Stack) */

	fmt.Println("Стек")
	fmt.Println("----")
	fmt.Println()

	stack := arraystack.New() // На основе массива
	// stack := linkedliststack.New() // На основе связного списка

	fmt.Println("Добавление элементов")
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
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
	value, ok = stack.Pop()
	fmt.Println("Pop:", value, ok)

	fmt.Println("Полная очистка")
	stack.Push(3)
	fmt.Printf("размер: %d, пуст: %t\n", stack.Size(), stack.Empty())
	stack.Clear()
	fmt.Printf("размер: %d, пуст: %t\n", stack.Size(), stack.Empty())
	fmt.Println()

	/* Карта (Map) */

	fmt.Println("Карта")
	fmt.Println("-----")
	fmt.Println()

	m := hashmap.New() // На основе хеш-таблицы (случайный порядок)
	// m := treemap.NewWithIntComparator() // На основе дерева (упорядочено)
	// m := linkedhashmap.New() // На основе хеш-таблицы и списка (в порядке вставки)
	// m := hashbidimap.New() // На основе хеш-таблицы (двусторонняя карта, случайный порядок)
	// m := treebidimap.NewWith( // На основе дерева (двусторонняя карта, упорядочено)
	// 	utils.IntComparator, utils.StringComparator)

	fmt.Println("Добавление элементов")
	m.Put(2, "b")
	m.Put(3, "c")
	m.Put(1, "x")
	m.Put(1, "a")
	// Переопределение ключа значением
	// (только для двусторонних карт)
	// m.Put(3, "a")
	b, _ := m.ToJSON()
	fmt.Println(string(b))

	fmt.Println("Получение по ключу")
	value, ok = m.Get(2)
	fmt.Println("2:", value, ok)
	value, ok = m.Get(4)
	fmt.Println("4:", value, ok)

	// Только для двусторонних карт
	// fmt.Println("Получение по значению")
	// value, ok = m.GetKey("a")
	// fmt.Println("a:", value, ok)
	// value, ok = m.GetKey("c")
	// fmt.Println("c:", value, ok)

	fmt.Println("Ключи и значения")
	fmt.Println(m.Keys(), m.Values())

	fmt.Println("Удаление элементов")
	m.Remove(3)
	fmt.Println(m.Values())

	fmt.Println("Полная очистка")
	fmt.Printf("размер: %d, пуста: %t\n", m.Size(), m.Empty())
	m.Clear()
	fmt.Printf("размер: %d, пуста: %t\n", m.Size(), m.Empty())
	fmt.Println()

	/* Дерево (Tree) */

	fmt.Println("Дерево")
	fmt.Println("------")
	fmt.Println()

	tree := redblacktree.NewWithIntComparator() // Красно-черное дерево
	// tree := avltree.NewWithIntComparator() // АВЛ-дерево
	// tree := btree.NewWithIntComparator(3) // B-дерево (BTree)

	fmt.Println("Добавление элементов")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a")
	tree.Put(3, "c")
	b, _ = tree.ToJSON()
	fmt.Println(string(b), "...")
	tree.Put(4, "d")
	tree.Put(5, "e")
	tree.Put(6, "f")
	// Еще один элемент для наглядности
	// (только для B-дерева)
	// tree.Put(7, "g")

	fmt.Println("Получение по ключу")
	value, ok = tree.Get(3)
	fmt.Println("3:", value, ok)
	value, ok = tree.Get(7)
	fmt.Println("7:", value, ok)

	fmt.Println("Ключи и значения")
	fmt.Println(tree.Keys(), tree.Values())

	fmt.Print(tree)

	fmt.Println("Удаление элементов")
	tree.Remove(2)

	fmt.Print(tree)

	fmt.Println("Дополнительные свойства")
	// Только для черно-красных и АВЛ-деревьев
	fmt.Println("мин. ключ:", tree.Left())
	fmt.Println("макс. ключ:", tree.Right())
	// Только для B-дерева
	// fmt.Println("высота:", tree.Height())
	// fmt.Println("мин. ключ:", tree.LeftKey())
	// fmt.Println("макс. ключ:", tree.RightKey())
	// fmt.Println("мин. значение:", tree.LeftValue())
	// fmt.Println("макс. значение:", tree.RightValue())

	fmt.Println("Полная очистка")
	fmt.Printf("размер: %d, пусто: %t\n", tree.Size(), tree.Empty())
	tree.Clear()
	fmt.Printf("размер: %d, пусто: %t\n", tree.Size(), tree.Empty())
	fmt.Println()

	/* Куча (Heap) */

	fmt.Println("Куча")
	fmt.Println("----")
	fmt.Println()

	// На основе дерева
	heap := binaryheap.NewWithIntComparator() // Минимальная куча
	// heap := binaryheap.NewWith(func(a, b interface{}) int { // Максимальная куча
	// 	return -utils.IntComparator(a, b)
	// })

	fmt.Println("Добавление элементов")
	heap.Push(2, 3)
	heap.Push(1)
	fmt.Println(heap.Values())

	fmt.Println("Просмотр и извлечение")
	value, ok = heap.Peek()
	fmt.Println("Peek:", value, ok)
	value, ok = heap.Pop()
	fmt.Println("Pop:", value, ok)
	value, ok = heap.Pop()
	fmt.Println("Pop:", value, ok)
	value, ok = heap.Pop()
	fmt.Println("Pop:", value, ok)
	value, ok = heap.Pop()
	fmt.Println("Pop:", value, ok)

	fmt.Println("Полная очистка")
	heap.Push(1)
	fmt.Printf("размер: %d, пуста: %t\n", heap.Size(), heap.Empty())
	heap.Clear()
	fmt.Printf("размер: %d, пуста: %t\n", heap.Size(), heap.Empty())
	fmt.Println()

	/* Очередь (Queue) */

	fmt.Println("Очередь")
	fmt.Println("-------")
	fmt.Println()

	// queue := arrayqueue.New() // На основе массива
	// queue := linkedlistqueue.New() // На основе связного списка
	queue := circularbuffer.New(3) // Кольцевой буфер

	fmt.Println("Добавление элементов")
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.Values())
	// Еще один элемент для наглядности
	// (только для кольцевого буфера)
	queue.Enqueue(4)
	fmt.Println(queue.Values())

	fmt.Println("Просмотр и извлечение")
	value, ok = queue.Peek()
	fmt.Println("Peek:", value, ok)
	value, ok = queue.Dequeue()
	fmt.Println("Dequeue:", value, ok)
	value, ok = queue.Dequeue()
	fmt.Println("Dequeue:", value, ok)
	value, ok = queue.Dequeue()
	fmt.Println("Dequeue:", value, ok)
	value, ok = queue.Dequeue()
	fmt.Println("Dequeue:", value, ok)

	fmt.Println("Полная очистка")
	queue.Enqueue(3)
	fmt.Printf("размер: %d, пуста: %t\n", queue.Size(), queue.Empty())
	queue.Clear()
	fmt.Printf("размер: %d, пуста: %t\n", queue.Size(), queue.Empty())
}
