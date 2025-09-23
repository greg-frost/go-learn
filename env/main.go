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
	fmt.Printf("FOO = %q\n", os.Getenv("FOO"))
	fmt.Printf("BAR = %q", os.Getenv("BAR"))
	if _, ok := os.LookupEnv("BAR"); !ok {
		fmt.Print(" (не задано)")
	}
	fmt.Println()
	fmt.Println()

	// Список всех переменных
	for _, e := range os.Environ()[:15] {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
	fmt.Println("...")
}
