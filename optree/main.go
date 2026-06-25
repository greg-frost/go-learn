package main

import (
	"errors"
	"fmt"
)

// Структура "узел дерева"
type TreeNode struct {
	val   TreeVal
	left  *TreeNode
	right *TreeNode
}

// Интерфейс "значение узла дерева"
type TreeVal interface {
	IsToken()
}

// Тип "число"
type Number int

func (Number) IsToken() {}

// Тип "оператор"
type Operator func(int, int) int

func (Operator) IsToken() {}

// Выполнение операции
func (o Operator) process(n1, n2 int) int {
	return o(n1, n2)
}

// Карта операций
var operators = map[string]Operator{
	"+": func(n1, n2 int) int {
		return n1 + n2
	},
	"-": func(n1, n2 int) int {
		return n1 - n2
	},
	"*": func(n1, n2 int) int {
		return n1 * n2
	},
	"/": func(n1, n2 int) int {
		return n1 / n2
	},
}

// Обход дерева
func WalkTree(t *TreeNode) (int, error) {
	switch val := t.val.(type) {
	case nil:
		return 0, errors.New("Неправильное выражение")
	case Number:
		return int(val), nil
	case Operator:
		left, err := WalkTree(t.left)
		if err != nil {
			return 0, err
		}
		right, err := WalkTree(t.right)
		if err != nil {
			return 0, err
		}
		return val.process(left, right), nil
	default:
		return 0, errors.New("Неизвестный тип узла")
	}
}

// Печать дерева
func (t *TreeNode) String() string {
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
	fmt.Println(" \n[ ДЕРЕВО ОПЕРАЦИЙ ]\n ")

	// Дерево
	tree := &TreeNode{
		val: operators["+"],
		left: &TreeNode{
			val:   operators["*"],
			left:  &TreeNode{val: Number(5)},
			right: &TreeNode{val: Number(10)},
		},
		right: &TreeNode{val: Number(20)},
	}

	// Печать
	fmt.Println("Дерево:")
	fmt.Println(tree)
	fmt.Println()

	// Обход
	fmt.Println("Выражение:")
	result, _ := WalkTree(tree)
	fmt.Println("5 * 10 + 20 =", result)
}
