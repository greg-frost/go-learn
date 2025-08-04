package main

import (
	"fmt"
	"strings"
)

// Интерфейс "итератор"
type Iterator interface {
	Next() string
	HasNext() bool
}

// Структура "массив"
type Array struct {
	values []string
}

// Конструктор массива
func NewArray() *Array {
	return new(Array)
}

// Добавление элементов
func (a *Array) Add(values ...string) {
	a.values = append(a.values, values...)
}

// Получение итератора для массива
func (a *Array) GetIterator() Iterator {
	return &ArrayIterator{
		values: a.values,
	}
}

// Структура "итератор массива"
type ArrayIterator struct {
	values []string
	pos    int
}

// Получение следующего элемента
func (i *ArrayIterator) Next() string {
	if !i.HasNext() {
		return ""
	}
	value := i.values[i.pos]
	i.pos++
	return value
}

// Проверка существования следующего элемента
func (i *ArrayIterator) HasNext() bool {
	return i.pos < len(i.values)
}

// Структура "список"
type List struct {
	head *Node
}

// Структура "узел"
type Node struct {
	value string
	next  *Node
}

// Конструктор списка
func NewList() *List {
	return &List{
		head: new(Node),
	}
}

// Добавление элементов
func (l *List) Add(values ...string) {
	curr := l.head
	for _, value := range values {
		curr.next = &Node{
			value: value,
		}
		curr = curr.next
	}
}

// Получение итератора для списка
func (l *List) GetIterator() Iterator {
	return &ListIterator{
		curr: l.head.next,
	}
}

// Структура "итератор списка"
type ListIterator struct {
	curr *Node
}

// Получение следующего элемента
func (i *ListIterator) Next() string {
	if !i.HasNext() {
		return ""
	}
	value := i.curr.value
	i.curr = i.curr.next
	return value
}

// Проверка существования следующего элемента
func (i *ListIterator) HasNext() bool {
	return i.curr != nil
}

// Структура "хеш-таблица"
type Hash struct {
	values map[string]string
}

// Конструктор хеш-таблицы
func NewHash() *Hash {
	return &Hash{
		values: make(map[string]string),
	}
}

// Добавление элементов
func (a *Hash) Add(values ...string) {
	for _, value := range values {
		a.values[strings.ToLower(value)] = value
	}
}

// Получение итератора для хеш-таблицы
func (h *Hash) GetIterator() Iterator {
	values := make([]string, 0, len(h.values))
	for _, value := range h.values {
		values = append(values, value)
	}
	return &HashIterator{
		values: values,
	}
}

// Структура "итератор хеш-таблицы"
type HashIterator struct {
	values []string
	pos    int
}

// Получение следующего элемента
func (i *HashIterator) Next() string {
	if !i.HasNext() {
		return ""
	}
	value := i.values[i.pos]
	i.pos++
	return value
}

// Проверка существования следующего элемента
func (i *HashIterator) HasNext() bool {
	return i.pos < len(i.values)
}

// Печать итератора
func Print(it Iterator) {
	if it.HasNext() {
		for it.HasNext() {
			fmt.Print(it.Next(), " ")
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println(" \n[ ИТЕРАТОР ]\n ")

	// Итератор массива
	array := NewArray()
	array.Add("Hello", "Hell", "World")
	fmt.Println("Массив:")
	iterator := array.GetIterator()
	Print(iterator) // Будет напечатано
	Print(iterator) // Будет проигнорировано
	fmt.Println()

	// Итератор списка
	list := NewList()
	list.Add("Goodbye", "Cruel", "World")
	fmt.Println("Список:")
	Print(list.GetIterator())
	fmt.Println()

	// Итератор хеш-таблицы
	hash := NewHash()
	hash.Add("Random", "Hashed", "Strings")
	fmt.Println("Хеш-таблица:")
	Print(hash.GetIterator())
}
