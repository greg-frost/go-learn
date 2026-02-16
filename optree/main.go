package main

import (
	"errors"
	"fmt"
)

// Структура "узел дерева"
type treeNode struct {
	val   treeVal
	left  *treeNode
	right *treeNode
}

// Интерфейс "значение узла дерева"
type treeVal interface {
	isToken()
}

// Тип "число"
type number int

func (number) isToken() {}

// Тип "оператор"
type operator func(int, int) int

func (operator) isToken() {}

// Выполнение операции
func (o operator) process(n1, n2 int) int {
	return o(n1, n2)
}

// Карта операций
var operators = map[string]operator{
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
func walkTree(t *treeNode) (int, error) {
	switch val := t.val.(type) {
	case nil:
		return 0, errors.New("Неправильное выражение")
	case number:
		return int(val), nil
	case operator:
		left, err := walkTree(t.left)
		if err != nil {
			return 0, err
		}
		right, err := walkTree(t.right)
		if err != nil {
			return 0, err
		}
		return val.process(left, right), nil
	default:
		return 0, errors.New("Неизвестный тип узла")
	}
}

// Печать дерева
func (t *treeNode) String() string {
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
	tree := &treeNode{
		val: operators["+"],
		left: &treeNode{
			val:   operators["*"],
			left:  &treeNode{val: number(5)},
			right: &treeNode{val: number(10)},
		},
		right: &treeNode{val: number(20)},
	}

	// Печать
	fmt.Println("Дерево:")
	fmt.Println(tree)
	fmt.Println()

	// Обход
	fmt.Println("Выражение:")
	result, _ := walkTree(tree)
	fmt.Println("5 * 10 + 20 =", result)
}
