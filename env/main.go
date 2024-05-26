package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(" \n[ ПЕРЕМЕННЫЕ ОКРУЖЕНИЯ ]\n ")

	// Установка
	os.Setenv("FOO", "1")

	// Чтение
	fmt.Println("FOO", os.Getenv("FOO"))
	fmt.Println("BAR", os.Getenv("BAR"))
	fmt.Println()

	// Список всех переменных
	for _, e := range os.Environ()[:15] {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
	fmt.Println("...")
}
