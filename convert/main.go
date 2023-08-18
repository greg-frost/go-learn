package main

import (
	"fmt"
)

func main() {
	fmt.Println(" \n[ КОНВЕРТЕР ]\n ")

	var f float32

	/* Температура */

	fmt.Print("Введите температуру по Фаренгейту: ")
	fmt.Scanf("%f", &f)

	fmt.Println(f, "по Фаренгейту равна", fmt.Sprintf("%.2f", (float32(f-32)*5/9)), "по Цельсию")

	fmt.Println()

	/* Расстояние */

	fmt.Print("Введите расстояние в футах: ")
	fmt.Scanf("%f", &f)

	fmt.Printf("%.f футов равно %.2f метров", f, float32(f)*0.3048)

	fmt.Println()
}
