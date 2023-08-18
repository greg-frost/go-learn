package main

import (
	"fmt"
)

func main() {
	fmt.Println(" \n[ ЦИКЛЫ ]\n ")

	/* If-Case */

	for i := 1; i <= 100; i++ {
		if t := i % 3; t == 0 {

			/* Выбор (case) */

			switch {
			case i%15 == 0:
				fmt.Print("FizzBuzz ")
			case i%5 == 0:
				fmt.Print("Buzz ")
			case i%3 == 0:
				fmt.Print("Fizz ")
			default:
				fmt.Print(i, " ")
			}

		} else if t == 1 {

			/* Условия (вариант 1) */

			if i%3 == 0 && i%5 == 0 {
				fmt.Print("FizzBuzz ")
			} else if i%3 == 0 {
				fmt.Print("Fizz ")
			} else if i%5 == 0 {
				fmt.Print("Buzz ")
			} else {
				fmt.Print(i, " ")
			}

		} else if t == 2 {

			/* Условия (вариант 2) */

			var found = false
			if i%3 == 0 {
				fmt.Print("Fizz")
				found = true
			}
			if i%5 == 0 {
				fmt.Print("Buzz")
				found = true
			}
			if !found {
				fmt.Print(i)
			}
			fmt.Print(" ")

		}
	}

	/* Метки */

	fmt.Println(" \n ")

	fmt.Println("Метки:")
	fmt.Println()

outerBreakLabel:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("[%d, %d]\n", i, j)
			break outerBreakLabel
		}
	}
	fmt.Println("Брейк!")

	fmt.Println()

outerContinueLabel:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("[%d, %d]\n", i, j)
			continue outerContinueLabel
		}
	}
	fmt.Println("Континью!")
}
