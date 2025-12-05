package main

import (
	"fmt"
	"math/big"
)

// Число Фибоначчи
func fib(n int) *big.Int {
	if n < 2 {
		return big.NewInt(int64(n))
	}

	a, b := big.NewInt(0), big.NewInt(1)

	for n--; n > 0; n-- {
		a.Add(a, b)
		a, b = b, a
	}

	return b
}

func main() {
	fmt.Println(" \n[ BIG INT ]\n ")

	size := 1000
	fmt.Printf("Число Фибоначчи (%d):\n\n%v\n\n", size, fib(size))

	size = 5000
	fmt.Printf("Число Фибоначчи (%d):\n\n%v\n\n", size, fib(size))

	size = 10000
	fmt.Printf("Число Фибоначчи (%d):\n\n%v\n", size, fib(size))
}
