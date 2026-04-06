package main

import (
	"bufio"
	"fmt"
	"go-learn/base"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Сканирование строк в файле
// (лучше было бы передать os.File, чем имя файла)
func ScanLinesInFile(filename string) {
	// Открытие файла
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Сканирование
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		upper := strings.ToUpper(scanner.Text())
		fmt.Printf("%q\n", upper)
	}

	// Обработка ошибок
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Путь
var path = base.Dir("lines")

func main() {
	fmt.Println(" \n[ ЛИНЕЙНЫЕ ФИЛЬТРЫ ]\n ")

	// Файл
	fmt.Println("Из файла:")
	ScanLinesInFile(filepath.Join(path, "data.txt"))
	fmt.Println()

	// Stdin
	fmt.Println("Вводите строки:")
	fmt.Println("(или нажмите Ctrl+C)")
	fmt.Println()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		upper := strings.ToUpper(scanner.Text())
		fmt.Printf("%q\n", upper)
	}

	// Обработка ошибок
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
