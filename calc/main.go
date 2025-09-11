package main

import (
	"fmt"
	"strconv"
)

// Сложение
func add(i int, j int) int {
	return i + j
}

// Вычитание
func sub(i int, j int) int {
	return i - j
}

// Умножение
func mul(i int, j int) int {
	return i * j
}

// Деление
func div(i int, j int) int {
	return i / j
}

// Карта операций
var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	fmt.Println(" \n[ КАЛЬКУЛЯТОР ]\n ")

	// Выражения
	expressions := [][]string{
		{"10", "+", "5"},
		{"10", "-", "5"},
		{"10", "*", "5"},
		{"10", "/", "5"},
		{"10", "%", "5"},
		{"15"},
		{"ten", "+", "five"},
	}

	// Вычисление
	for _, expr := range expressions {
		if len(expr) != 3 {
			fmt.Println("Неправильное выражение:", expr)
			continue
		}

		// Левый операнд
		p1, err := strconv.Atoi(expr[0])
		if err != nil {
			fmt.Println("Ошибка левого операнда:", err)
			continue
		}

		// Операция
		op := expr[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("Неизвестная операция:", op)
			continue
		}

		// Правый операнд
		p2, err := strconv.Atoi(expr[2])
		if err != nil {
			fmt.Println("Ошибка правого операнда:", err)
			continue
		}

		// Результат
		result := opFunc(p1, p2)
		fmt.Println(p1, op, p2, "=", result)
	}
}
