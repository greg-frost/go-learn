package main

import (
	"fmt"
)

// Структура "дерево"
type Tree struct {
	x int
	y int
	r float64
}

// Конструктор дерева
func NewTree(x, y int, r float64) *Tree {
	return &Tree{
		x: x,
		y: y,
		r: r,
	}
}

// Прорисовка дерева
func (t *Tree) Draw() {
	fmt.Printf("Дерево: %.1f (%d, %d)\n", t.r, t.x, t.y)
}

// Структура "лес"
type Forest struct {
	trees []*Tree
}

// Конструктор леса
func NewForest(trees ...*Tree) *Forest {
	return &Forest{
		trees: trees,
	}
}

// Прорисовка леса
func (f *Forest) Draw() {
	for _, tree := range f.trees {
		tree.Draw()
	}
}

func main() {
	fmt.Println(" \n[ ПРИСПОСОБЛЕНЕЦ (ЛЕГКОВЕС) ]\n ")

	// Медленный лес (каждое дерево - объект)
	forest := NewForest(
		NewTree(1, 2, 3.5),
		NewTree(3, 0, 2.1),
		NewTree(5, 7, 6.2),
	)
	fmt.Println("Медленный лес:")
	forest.Draw()
}
