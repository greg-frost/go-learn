package main

import (
	"fmt"
)

// Структура "узел дерева"
type Tree struct {
	val         int
	left, right *Tree
}

// Создание дерева
func MakeTree(values []int) *Tree {
	var t *Tree
	for _, v := range values {
		t = t.Insert(v)
	}
	return t
}

// Добавление элемента
func (t *Tree) Insert(val int) *Tree {
	if t == nil {
		return &Tree{val: val}
	}
	if val < t.val {
		t.left = t.left.Insert(val)
	} else if val > t.val {
		t.right = t.right.Insert(val)
	}
	return t
}

// Поиск элемента
func (t *Tree) Contains(val int) bool {
	switch {
	case t == nil:
		return false
	case val < t.val:
		return t.left.Contains(val)
	case val > t.val:
		return t.right.Contains(val)
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
	s += fmt.Sprint(t.val)
	if t.right != nil {
		s += " " + t.right.String()
	}
	return "(" + s + ")"
}

func main() {
	fmt.Println(" \n[ ДЕРЕВО ]\n ")

	values := []int{10, 5, 25, 50, 100, 2, 3, 1}
	search := []int{25, 30}

	tree := MakeTree(values)

	fmt.Println("Дерево:")
	fmt.Println(tree)

	fmt.Println()

	for _, v := range search {
		fmt.Println("Есть ли", v, ":", tree.Contains(v))
	}
}
