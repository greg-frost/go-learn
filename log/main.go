package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(" \n[ ЛОГИРОВАНИЕ ]\n ")

	// Стандартный
	fmt.Println("Стандартный логгер:")
	log.Printf("Hello\n")
	log.Println("World")
	fmt.Println()

	// Настроенный
	fmt.Println("Настроенный логгер:")
	flags := log.LstdFlags | log.Lshortfile
	logger := log.New(os.Stderr, "[stderr] ", flags)
	logger.Printf("Hello\n")
	logger.Println("World")
	fmt.Println()

	// Файловый
	fmt.Println("Файловый логгер:")
	f, err := os.CreateTemp("", "temp")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name())
	file := log.New(f, "[file] ", flags)
	file.Printf("Hello\n")
	file.Println("World")
	bs, err := os.ReadFile(f.Name())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(bs))
}
