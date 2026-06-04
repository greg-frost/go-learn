package main

import (
	"fmt"
)

func main() {
	fmt.Println(" \n[ ТЕСТИРОВАНИЕ ]\n ")

	modules := []string{
		"math", "slice",
		"control", "flaky",
	}

	fmt.Println("Модули:")
	for _, module := range modules {
		fmt.Println(module)
	}
}
