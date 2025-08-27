package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Структура "дерево"
type Tree struct {
	x      int
	y      int
	r      int
	leaves []int
}

// Конструктор дерева
func NewTree(x, y, r int) *Tree {
	return &Tree{
		x:      x,
		y:      y,
		r:      r,
		leaves: randomLeaves(leavesSize),
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
	var size int
	for _, tree := range f.trees {
		size += len(tree.leaves)
	}
	return size
}

// Структура "легковес"
type Flyweight struct {
	trees  [][3]int
	leaves []int
}

// Конструктор легковеса
func NewFlyweight(trees ...[3]int) *Flyweight {
	treesData := make([][3]int, 0, len(trees))
	for _, tree := range trees {
		treesData = append(treesData, tree)
	}
	return &Flyweight{
		trees:  treesData,
		leaves: randomLeaves(leavesSize),
	}
}

// Прорисовка легковеса
func (f *Flyweight) Draw() {
	for i := 0; i < len(f.trees); i++ {
		drawTree(f.trees[i][0], f.trees[i][1], f.trees[i][2], f.leaves)
	}
}

// Размер легковеса
func (f *Flyweight) Size() int {
	return len(f.leaves)
}

// Размер массива листьев
const leavesSize = 1e6

// Генерация случайных листьев
func randomLeaves(size int) []int {
	leaves := make([]int, size)
	for i := 0; i < size; i++ {
		leaves[i] = rand.Intn(math.MaxInt)
	}
	return leaves
}

// Прорисовка абстрактного дерева
func drawTree(x, y, r int, leaves []int) {
	fmt.Printf("Дерево: %d м (%d, %d)\n", r, x, y)
}

func main() {
	fmt.Println(" \n[ ЛЕГКОВЕС (ПРИСПОСОБЛЕНЕЦ) ]\n ")

	// Медленный лес:
	// каждое дерево - объект,
	// свои листья у каждого дерева
	forest := NewForest(
		NewTree(1, 2, 3),
		NewTree(3, 0, 2),
		NewTree(5, 7, 6),
	)
	fmt.Println("Медленный лес:")
	forest.Draw()
	fmt.Printf("Размер: %.1f MB\n",
		float64(forest.Size())/1e6)
	fmt.Println()

	// Быстрый лес:
	// все деревья в одном объекте,
	// общие листья у всех деревьев
	flyweight := NewFlyweight(
		[3]int{1, 2, 3},
		[3]int{3, 0, 2},
		[3]int{5, 7, 6},
	)
	fmt.Println("Быстрый лес:")
	flyweight.Draw()
	fmt.Printf("Размер: %.1f MB\n",
		float64(flyweight.Size())/1e6)
}
