package main

import (
	"fmt"
)

// Структура "структура"
type Struct struct {
	state bool
}

// Печать по значению
func (s Struct) PrintByValue() {
	fmt.Println("По значению:", s.state)
}

// Печать по указателю
func (s *Struct) PrintByPointer() {
	fmt.Println("По указателю:", s.state)
}

func main() {
	fmt.Println(" \n[ DEFER ]\n ")

	// Вычисление аргументов
	var variable, closure bool
	defer func(local bool) {
		fmt.Println()
		fmt.Println("Вычисление аргументов")
		fmt.Println("Локальное:", local)
		fmt.Println("Замыкание:", closure)
	}(variable)
	variable = true // Не увидит изменение
	closure = true  // Увидит изменение

	// Значение и указатель
	fmt.Println("Значение и указатель")
	p := &Struct{}
	defer p.PrintByPointer()
	p.state = true
	v := Struct{}
	defer v.PrintByValue()
	v.state = true
}
