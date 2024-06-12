package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

// Внешняя функция
func outer() {
	inner()
}

// Внутренняя функция
func inner() {
	// Печать трассировки
	fmt.Println("Печать:")
	debug.PrintStack()
	fmt.Println()

	// Получение трассировки
	fmt.Println("Получение:")
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, false)
	fmt.Println(string(buf[:n]))
}

func main() {
	fmt.Println(" \n[ ТРАССИРОВКА ]\n ")

	// Вход
	outer()
}
