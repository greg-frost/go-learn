package main

import (
	"fmt"
)

const (
	i         = 10
	f         = 1.5
	i64 int64 = 88
)

var v = 45

func main() {
	fmt.Println(" \n[ КОНСТАНТЫ ]\n ")

	fmt.Println("i + f = ", i+f)
	fmt.Println("i + i64 = ", i+i64)
	fmt.Println("i + v = ", i+v)

	// f + v = f (untyped float constant 1.5) truncated to int
	// f + i64 = f (untyped float constant 1.5) truncated to int64
	// i64 + v = i64 + v (mismatched types int64 and int)

	//i = 20 // не выйдет
	//pi := &i // а вот это попробуйте
	//*pi = 20 // тоже фиаско
}
