package main

import (
	"fmt"
)

// Структура "массив"
type Array struct {
	values []string
}

// Конструктор массива
func NewArray(values []string) *Array {
	return &Array{
		values: values,
	}
}

// Получение итератора для массива
func (a *Array) GetIterator() Iterator {
	return &ArrayIterator{
		values: &a.values,
	}
}

// Интерфейс "итератор"
type Iterator interface {
	Next() string
	HasNext() bool
}

// Структура "итератор массива"
type ArrayIterator struct {
	values *[]string
	pos    int
}

// Получение следующего элемента
func (i *ArrayIterator) Next() string {
	if i.pos == len(*i.values) {
		return ""
	}
	next := (*i.values)[i.pos]
	i.pos++
	return next
}

// Проверка существования следующего элемента
func (i *ArrayIterator) HasNext() bool {
	return i.pos < len(*i.values)
}

// Печать итератора
func Print(it Iterator) {
	for it.HasNext() {
		value := it.Next()
		fmt.Print(value, " ")
	}
	fmt.Println()
}

func main() {
	fmt.Println(" \n[ ИТЕРАТОР ]\n ")

	// Итератор массива
	array := NewArray([]string{"Hello", "Hell", "World"})
	iterator := array.GetIterator()
	fmt.Println("Массив:")
	Print(iterator)
}
