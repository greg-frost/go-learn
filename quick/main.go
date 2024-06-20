package main

import (
	"fmt"
	"strings"
)

// Дополнение/обрезание строки
func Pad(s string, max uint) string {
	size := uint(len(s))
	if size > max {
		// return s[:max-1] // Ошибка!
		return s[:max]
	}
	return strings.Repeat(" ", int(max-size)) + s
}

func main() {
	fmt.Println(" \n[ ПОРОЖДАЮЩЕЕ ТЕСТИРОВАНИЕ ]\n ")

	var size uint = 4
	fmt.Printf("Обрезание:  %q [%d]\n", Pad("longer", size), size)
	fmt.Printf("Дополнение: %q [%d]\n", Pad("s", size), size)
}
