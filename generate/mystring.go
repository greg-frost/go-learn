//go:generate ./queue MyString
package main

import (
	"fmt"
)

// Пользовательский тип
type MyString string

func main() {
	fmt.Println(" \n[ КОДОГЕНЕРАЦИЯ (MYSTRING) ]\n ")

	var one, two, three MyString = "one", "two", "three"

	q := NewMyStringQueue()
	q.Insert(one)
	q.Insert(two)
	q.Insert(three)

	fmt.Println("Первое значение:", q.Remove())
}
