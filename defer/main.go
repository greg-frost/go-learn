package main

import (
	"fmt"
)

func main() {
	fmt.Println(" \n[ DEFER ]\n ")

	// Вычисление аргументов
	var variable, closure bool
	defer func(local bool) {
		fmt.Println("Вычисление аргументов")
		fmt.Println("Локальное:", local)
		fmt.Println("Замыкание:", closure)
	}(variable)
	variable = true // Не увидит изменение
	closure = true  // Увидит изменение
}
