//go:generate ./queue MyInt
package main

import (
	"fmt"
)

// Пользовательский тип
type MyInt int

func main() {
	fmt.Println(" \n[ КОДОГЕНЕРАЦИЯ (MYINT) ]\n ")

	var one, two, three MyInt = 1, 2, 3

	q := NewMyIntQueue()
	q.Insert(one)
	q.Insert(two)
	q.Insert(three)

	fmt.Println("Первое значение:", q.Remove())
}
