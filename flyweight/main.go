package main

import (
	"fmt"
)

// Размер массива листьев
const leavesSize = 1e6

// Структура "дерево"
type Tree struct {
	x      int
	y      int
	r      float64
	leaves []int
}

// Конструктор дерева
func NewTree(x, y int, r float64) *Tree {
	return &Tree{
		x:      x,
		y:      y,
		r:      r,
		leaves: make([]int, leavesSize),
	}
}

// Прорисовка дерева
func (t *Tree) Draw() {
	drawTree(t.x, t.y, t.r, t.leaves)
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

// Размер леса
func (f *Forest) Size() int {
	var res int
	for _, tree := range f.trees {
		res += len(tree.leaves)
	}
	return res
}

// Структура "легковес"
type Flyweight struct {
	coords [][2]int
	rads   []float64
	leaves []int
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
		leaves: make([]int, leavesSize),
	}
}

// Прорисовка легковеса
func (f *Flyweight) Draw() {
	for i := 0; i < len(f.coords); i++ {
		drawTree(f.coords[i][0], f.coords[i][1], f.rads[i], f.leaves)
	}
}

// Размер легковеса
func (f *Flyweight) Size() int {
	return len(f.leaves)
}

// Прорисовка абстрактного дерева
func drawTree(x, y int, r float64, leaves []int) {
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
	fmt.Printf("Размер: %.1f MB\n",
		float64(forest.Size())/1e6)
	fmt.Println()

	// Быстрый лес (все деревья в одном месте)
	flyweight := NewFlyweight(
		[3]interface{}{1, 2, 3.5},
		[3]interface{}{3, 0, 2.1},
		[3]interface{}{5, 7, 6.2},
	)
	fmt.Println("Быстрый лес:")
	flyweight.Draw()
	fmt.Printf("Размер: %.1f MB\n",
		float64(flyweight.Size())/1e6)
}
