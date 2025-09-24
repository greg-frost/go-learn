package main

import (
	"fmt"
)

func main() {
	fmt.Println(" \n[ КОНВЕРТЕР ]\n ")

	var f float32

	// Температура
	fmt.Print("Введите температуру по Фаренгейту: ")
	fmt.Scanln(&f)
	fmt.Printf("%.2f по Фаренгейту равна %.2f по Цельсию", f, float32(f-32)*5/9)
	fmt.Println()

	// Расстояние
	fmt.Print("Введите расстояние в футах: ")
	fmt.Scanln(&f)
	fmt.Printf("%.2f футов равно %.2f метров", f, float32(f)*0.3048)
	fmt.Println()
}
