package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println(" \n[ АРГУМЕНТЫ ]\n ")

	// Получение
	var intValue int
	flag.IntVar(&intValue, "num", 0, "числовой флаг")
	boolValue := flag.Bool("cond", false, "логический флаг")
	strValue := flag.String("action", "none", "флаг действия")

	// Парсинг
	flag.Parse()

	// Вывод флагов
	fmt.Println("Числовой флаг:", intValue)
	fmt.Println("Логический флаг:", *boolValue)
	fmt.Println("Флаг действия:", *strValue)

	// Вывод остальных аргументов
	fmt.Println("Аргументы:", flag.Args())
}
