package main

import (
	"container/list"
	"fmt"
)

// Структура "связный список"
type LinkedList struct {
	Value interface{}
	Next  *LinkedList
}

// Вставка элемента
func (ll *LinkedList) Insert(pos int, val interface{}) *LinkedList {
	if ll == nil || pos == 0 {
		return &LinkedList{
			Value: val,
			Next:  ll,
		}
	}
	fmt.Print(".")
	ll.Next = ll.Next.Insert(pos-1, val)
	return ll
}

// Печать списка
func (ll *LinkedList) String() string {
	if ll == nil {
		return "(nil)"
	}
	return fmt.Sprintf("%v > %s", ll.Value, ll.Next)
}

func main() {
	fmt.Println(" \n[ СПИСКИ ]\n ")

	// Библиотечный список
	fmt.Println("Библиотечный")
	fmt.Println("------------")
	var l list.List
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.PushFront(2)
	l.PushFront(1)
	l.PushFront(0)
	fmt.Print("Вперед: ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
	fmt.Print("Назад: ")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
	fmt.Println()

	// Собственный список
	fmt.Println("Собственный")
	fmt.Println("-----------")
	var ll *LinkedList
	ll = ll.Insert(0, 10)
	fmt.Println(ll)
	ll = ll.Insert(0, "foo")
	fmt.Println(ll)
	ll = ll.Insert(1, 4.5)
	fmt.Println(ll)
	ll = ll.Insert(5, false)
	fmt.Println(ll)
	ll = ll.Insert(2, struct{}{})
	fmt.Println(ll)
}
