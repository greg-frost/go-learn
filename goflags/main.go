package main

import (
	"fmt"

	flags "github.com/jessevdk/go-flags"
)

// Структура "параметры"
var opts struct {
	Name    string `short:"n" long:"name" default:"Greg" description:"Имя для приветствия"`
	English bool   `short:"e" long:"english" description:"Язык приветствия"`
}

func main() {
	fmt.Println(" \n[ GO-FLAGS ]\n ")

	// Парсинг
	flags.Parse(&opts)

	// Использование
	if opts.English {
		fmt.Printf("Hello, %s!\n", opts.Name)
	} else {
		fmt.Printf("Привет, %s!\n", opts.Name)
	}
}
