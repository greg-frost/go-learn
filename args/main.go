package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println(" \n[ АРГУМЕНТЫ ]\n ")

	/* Получение */

	flgInt := flag.Int("num", 0, "числовой флаг")
	flgStr := flag.String("action", "none", "флаг действия")
	flag.Parse()

	/* Вывод */

	fmt.Println("Числовой флаг:", *flgInt)
	fmt.Println("Флаг действия:", *flgStr)
	fmt.Println("Аргументы:", flag.Args())
}
