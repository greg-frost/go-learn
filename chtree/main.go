package main

import (
	"fmt"
	"math/rand"
)

// Структура "узел дерева"
type Tree struct {
	val         int
	left, right *Tree
}

// Создание дерева
func MakeTree(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = Insert(t, (1+v)*k)
	}
	return t
}

// Добавление элемента
func Insert(t *Tree, val int) *Tree {
	if t == nil {
		return &Tree{val: val}
	}
	if val < t.val {
		t.left = Insert(t.left, val)
	} else if val > t.val {
		t.right = Insert(t.right, val)
	}
	return t
}

// Обход дерева
func Walk(t *Tree, ch chan int) {
	if t.left != nil {
		Walk(t.left, ch)
	}

	ch <- t.val

	if t.right != nil {
		Walk(t.right, ch)
	}
}

// Сравнение деревьев
func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	var isSame = true
	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			isSame = false
		}
	}

	return isSame
}

// Печать дерева
func Print(t *Tree, ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Print(<-ch, " ")
	}
	fmt.Println()
}

func main() {
	fmt.Println(" \n[ ДЕРЕВО-КАНАЛ ]\n ")

	ch := make(chan int)

	/* Создание и печать деревьев */

	fmt.Println("Первое дерево:")
	t1 := MakeTree(1)
	go Walk(t1, ch)
	Print(t1, ch)

	fmt.Println()

	fmt.Println("Второе дерево:")
	t2 := MakeTree(2)
	go Walk(t2, ch)
	Print(t2, ch)

	fmt.Println()

	/* Сравнение деревьев */

	t3 := t1
	fmt.Println("Третье дерево:\n(такое же, как первое)")

	fmt.Println()

	fmt.Println("Равенство деревьев:")
	fmt.Println("1 и 2:", Same(t1, t2))
	fmt.Println("1 и 3:", Same(t1, t3))
}
