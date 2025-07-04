package main

import (
	"fmt"
)

func main() {
	fmt.Println(" \n[ ЛОГИКА ]\n ")

	/* Простые выражения */

	//fmt.Println("True and True =", true && true)
	fmt.Println("True and False =", true && false)
	//fmt.Println("True or True =", true || true)
	fmt.Println("True or False =", true || false)
	fmt.Println("Not True =", !true)

	fmt.Println()

	/* Составное выражение */

	fmt.Println(
		"(True && False) or \n(False and True) or \nNot (True and False) =",
		(true && false) || (false && true) || !(true && false),
	)
}
