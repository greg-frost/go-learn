package main

import (
	"fmt"
	"math/rand"
)

// Структура "узел дерева"
type TreeNode struct {
	val         int
	left, right *TreeNode
}

// Конструктор дерева
func NewTree(k int) *TreeNode {
	var t *TreeNode
	for _, v := range rand.Perm(10) {
		t = InsertNode(t, (1+v)*k)
	}
	return t
}

// Добавление элемента
func InsertNode(t *TreeNode, val int) *TreeNode {
	if t == nil {
		return &TreeNode{val: val}
	}
	if val < t.val {
		t.left = InsertNode(t.left, val)
	} else if val > t.val {
		t.right = InsertNode(t.right, val)
	}
	return t
}

// Обход дерева
func WalkTree(t *TreeNode, ch chan int) {
	if t.left != nil {
		WalkTree(t.left, ch)
	}
	ch <- t.val
	if t.right != nil {
		WalkTree(t.right, ch)
	}
}

// Сравнение деревьев
func IsSame(t1, t2 *TreeNode) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go WalkTree(t1, ch1)
	go WalkTree(t2, ch2)

	var isSame = true
	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			isSame = false
		}
	}
	return isSame
}

// Печать дерева
func PrintTree(t *TreeNode, ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Print(<-ch, " ")
	}
	fmt.Println()
}

func main() {
	fmt.Println(" \n[ ДЕРЕВО-КАНАЛ ]\n ")

	// Канал
	ch := make(chan int)

	// Создание и печать деревьев
	fmt.Println("Первое дерево:")
	t1 := NewTree(1)
	go WalkTree(t1, ch)
	PrintTree(t1, ch)
	fmt.Println()

	fmt.Println("Второе дерево:")
	t2 := NewTree(2)
	go WalkTree(t2, ch)
	PrintTree(t2, ch)
	fmt.Println()

	// Сравнение деревьев
	t3 := t1
	fmt.Println("Третье дерево:\n(такое же, как первое)")
	fmt.Println()
	fmt.Println("Равенство деревьев:")
	fmt.Println("1 и 2:", IsSame(t1, t2))
	fmt.Println("1 и 3:", IsSame(t1, t3))
}
