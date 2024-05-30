package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(" \n[ ЛИНЕЙНЫЕ ФИЛЬТРЫ ]\n ")

	fmt.Println("Вводите строки:")
	fmt.Println("(или нажмите Ctrl+C)")
	fmt.Println()

	// Построчный стандартный ввод
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Преобразование в верхний регистр
		upper := strings.ToUpper(scanner.Text())
		fmt.Println(upper)
	}

	// Обработка ошибок
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка:", err)
		os.Exit(1)
	}
}
