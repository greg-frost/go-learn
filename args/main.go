package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println(" \n[ АРГУМЕНТЫ ]\n ")

	// Получение
	var flgInt int
	flag.IntVar(&flgInt, "num", 0, "числовой флаг")
	flgBool := flag.Bool("cond", false, "логический флаг")
	flgStr := flag.String("action", "none", "флаг действия")

	// Парсинг
	flag.Parse()

	// Вывод флагов
	fmt.Println("Числовой флаг:", flgInt)
	fmt.Println("Логический флаг:", *flgBool)
	fmt.Println("Флаг действия:", *flgStr)

	// Вывод остальных аргументов
	fmt.Println("Аргументы:", flag.Args())
}
