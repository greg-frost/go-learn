package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	fmt.Println(" \n[ OS ]\n ")

	// ОС
	fmt.Print("Операционная система: ")
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("WINDOWS")
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("LINUX")
	default:
		fmt.Println(strings.ToUpper(os))
	}

	// Число ядер
	fmt.Println("Число ядер процессора:", runtime.NumCPU())
}
