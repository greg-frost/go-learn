package main

import (
	"fmt"
)

// Структура "узел дерева"
type Tree struct {
	value       int
	left, right *Tree
}

// Создание дерева
func MakeTree(values []int) *Tree {
	var t *Tree
	for _, value := range values {
		t = t.Insert(value)
	}
	return t
}

// Добавление элемента
func (t *Tree) Insert(value int) *Tree {
	if t == nil {
		return &Tree{value: value}
	}
	if value < t.value {
		t.left = t.left.Insert(value)
	} else if value > t.value {
		t.right = t.right.Insert(value)
	}
	return t
}

// Поиск элемента
func (t *Tree) Contains(value int) bool {
	switch {
	case t == nil:
		return false
	case value < t.value:
		return t.left.Contains(value)
	case value > t.value:
		return t.right.Contains(value)
	default:
		return true
	}
}

// Печать дерева
func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.left != nil {
		s += t.left.String() + " "
	}
	s += fmt.Sprint(t.value)
	if t.right != nil {
		s += " " + t.right.String()
	}
	return "(" + s + ")"
}

func main() {
	fmt.Println(" \n[ ДЕРЕВО ]\n ")

	// Создание
	values := []int{10, 5, 25, 50, 100, 2, 3, 1}
	search := []int{25, 30}
	tree := MakeTree(values)

	// Печать
	fmt.Println("Дерево:")
	fmt.Println(tree)
	fmt.Println()

	// Поиск
	for _, value := range search {
		fmt.Printf("Поиск %d: %t\n", value, tree.Contains(value))
	}
}
