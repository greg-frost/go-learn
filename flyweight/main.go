package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

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
		leaves: randomLeaves(leavesSize),
	}
}

// Прорисовка дерева
func (t *Tree) Draw() {
	fmt.Printf("Дерево %.1f м (%d, %d)\n", t.r, t.x, t.y)
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

// Структура "лес-легковес"
type FlyweightForest struct {
	trees  []*FlyweightTree
	leaves []int
}

// Конструктор леса-легковеса
func NewFlyweightForest(trees ...*FlyweightTree) *FlyweightForest {
	return &FlyweightForest{
		trees:  trees,
		leaves: randomLeaves(leavesSize),
	}
}

// Прорисовка леса-легковеса
func (ff *FlyweightForest) Draw() {
	for _, tree := range ff.trees {
		tree.Draw(ff.leaves)
	}
}

// Размер леса-легковеса
func (ff *FlyweightForest) Size() int {
	return len(ff.leaves)
}

// Структура "дерево-легковес"
type FlyweightTree struct {
	x int
	y int
	r float64
}

// Конструктор дерева-легковеса
func NewFlyweightTree(x, y int, r float64) *FlyweightTree {
	return &FlyweightTree{
		x: x,
		y: y,
		r: r,
	}
}

// Прорисовка дерева-легковеса
func (ft *FlyweightTree) Draw(leaves []int) {
	fmt.Printf("Дерево %.1f м (%d, %d)\n", ft.r, ft.x, ft.y)
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

func main() {
	fmt.Println(" \n[ ЛЕГКОВЕС (ПРИСПОСОБЛЕНЕЦ) ]\n ")

	// Обычный лес
	// (свои листья у каждого дерева)
	start := time.Now()
	forest := NewForest(
		NewTree(1, 2, 3.5),
		NewTree(3, 0, 2.1),
		NewTree(5, 7, 6.2),
	)
	fmt.Println("Тяжелый и медленный лес:")
	forest.Draw()
	fmt.Println("Время:", time.Since(start))
	fmt.Printf("Размер: %.1f MB\n",
		float64(forest.Size())/1e6)
	fmt.Println()

	// Лес-легковес
	// (общие листья у всех деревьев)
	start = time.Now()
	flyweight := NewFlyweightForest(
		NewFlyweightTree(0, 0, 1.0),
		NewFlyweightTree(2, 3, 7.7),
		NewFlyweightTree(-1, -1, 5.3),
	)
	fmt.Println("Легкий и быстрый лес:")
	flyweight.Draw()
	fmt.Println("Время:", time.Since(start))
	fmt.Printf("Размер: %.1f MB\n",
		float64(flyweight.Size())/1e6)
}
