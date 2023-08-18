package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(" \n[ КОД ]\n ")

	files := 0
	lines := 0

	path := os.Getenv("GOPATH") + "/src/golearn/"

	/* Рекурсивный обход всех файлов .go */

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".go") {
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
	fmt.Println("Файлов .go:", files)
	fmt.Println("Строк кода:", lines)
}
