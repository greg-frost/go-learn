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
	drawTree(t.r, t.x, t.y)
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

// Структура "легковес"
type Flyweight struct {
	coords [][2]int
	rads   []float64
}

// Конструктор легковеса
func NewFlyweight(trees ...[3]interface{}) *Flyweight {
	coords := make([][2]int, 0, len(trees))
	rads := make([]float64, 0, len(trees))
	for _, tree := range trees {
		coords = append(coords, [2]int{
			tree[0].(int),
			tree[1].(int),
		})
		rads = append(rads, tree[2].(float64))
	}
	return &Flyweight{
		coords: coords,
		rads:   rads,
	}
}

// Прорисовка легковеса
func (f *Flyweight) Draw() {
	for i := 0; i < len(f.coords); i++ {
		drawTree(f.rads[i], f.coords[i][0], f.coords[i][1])
	}
}

// Прорисовка некого дерева
func drawTree(r float64, x, y int) {
	fmt.Printf("Дерево: %.1f (%d, %d)\n", r, x, y)
}

func main() {
	fmt.Println(" \n[ ЛЕГКОВЕС (ПРИСПОСОБЛЕНЕЦ) ]\n ")

	// Медленный лес (каждое дерево - объект)
	forest := NewForest(
		NewTree(1, 2, 3.5),
		NewTree(3, 0, 2.1),
		NewTree(5, 7, 6.2),
	)
	fmt.Println("Медленный лес:")
	forest.Draw()
	fmt.Println()

	// Быстрый лес (все деревья в одном месте)
	flyweight := NewFlyweight(
		[3]interface{}{1, 2, 3.5},
		[3]interface{}{3, 0, 2.1},
		[3]interface{}{5, 7, 6.2},
	)
	fmt.Println("Быстрый лес:")
	flyweight.Draw()
}
