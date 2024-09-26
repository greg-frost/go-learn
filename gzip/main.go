package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"

	"go-learn/base"
)

func main() {
	fmt.Println(" \n[ GZIP-КОМПРЕССИЯ ]\n ")

	// Смена директории
	path := base.Dir("gzip")
	os.Chdir(path)

	if len(os.Args) == 1 {
		fmt.Println("Передайте список файлов в виде параметров!")
		return
	}

	fmt.Println("Идет gzip-компрессия...")

	// Параллельное сжатие
	var wg sync.WaitGroup
	var count int
	for _, file := range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			if err := compress(filename); err != nil {
				fmt.Println("Ошибка:", err)
				return
			}
			count++
		}(file)
	}
	wg.Wait()

	fmt.Printf("Обработано файлов: %d\n", count)
}

// Сжатие файла
func compress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	defer gzout.Close()

	return err
}
