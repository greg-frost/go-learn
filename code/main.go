package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(" \n[ КОД ]\n ")

	// Настройки
	path := os.Getenv("GOPATH") + "/src/learn/"
	skip := []string{".pb", "generated"}

	var files, dirs, lines int

	/* Подсчет корневых директорий */

	root, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	rootFiles, err := root.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range rootFiles {
		if file.IsDir() && !strings.HasPrefix(file.Name(), ".") {
			dirs++
		}
	}

	/* Подсчет количества go-файлов и строк кода */

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".go") {
			for _, s := range skip {
				if strings.Contains(path, s) {
					return nil
				}
			}

			file, err := os.Open(path)
			if err != nil {
				return nil
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				lines++
			}

			files++
		}
		return nil
	})

	/* Вывод статистики */

	fmt.Println("Статистика:")
	fmt.Println()
	fmt.Println("Проектов:  ", dirs)
	fmt.Println("Файлов go: ", files)
	fmt.Println("Строк кода:", lines)
}
