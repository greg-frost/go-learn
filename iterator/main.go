package main

import (
	"fmt"
)

// Структура "массив"
type Array struct {
	values []string
}

// Конструктор массива
func NewArray() *Array {
	return new(Array)
}

// Добавление элементов
func (a *Array) Add(elements ...string) {
	for _, value := range elements {
		a.values = append(a.values, value)
	}
}

// Получение итератора для массива
func (a *Array) GetIterator() Iterator {
	return &ArrayIterator{
		values: a.values,
	}
}

// Интерфейс "итератор"
type Iterator interface {
	Next() string
	HasNext() bool
}

// Структура "итератор массива"
type ArrayIterator struct {
	values []string
	pos    int
}

// Получение следующего элемента
func (i *ArrayIterator) Next() string {
	if i.pos == len(i.values) {
		return ""
	}
	next := (i.values)[i.pos]
	i.pos++
	return next
}

// Проверка существования следующего элемента
func (i *ArrayIterator) HasNext() bool {
	return i.pos < len(i.values)
}

// Печать итератора
func Print(it Iterator) {
	if it.HasNext() {
		for it.HasNext() {
			value := it.Next()
			fmt.Print(value, " ")
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
}
