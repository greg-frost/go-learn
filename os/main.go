package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	fmt.Println(" \n[ OS ]\n ")

	fmt.Print("Go работает на ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("LINUX.")
	default:
		fmt.Printf("%s.\n", strings.ToUpper(os))
	}
	fmt.Println("Число ядер процессора:", runtime.NumCPU())
}
