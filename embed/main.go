package main

import (
	"embed"
	"fmt"
)

//go:embed data/hello.txt
var fileStringHello string

//go:embed data/welcome.txt
var fileStringWelcome []byte

//go:embed data
var folder embed.FS

func main() {
	fmt.Println(" \n[ ВСТРАИВАНИЕ ]\n ")

	// Файлы
	fmt.Println("Файл hello.txt:", fileStringHello)
	fmt.Println("Файл welcome.txt:", string(fileStringWelcome))
	fmt.Println()

	// Папка
	fmt.Println("Директория data:")
	dir, _ := folder.ReadDir("data")
	for _, d := range dir {
		i, _ := d.Info()
		fmt.Printf("%s (%d байт)\n", d.Name(), i.Size())
	}
}
