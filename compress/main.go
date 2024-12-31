package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"go-learn/base"
)

func main() {
	fmt.Println(" \n[ СЖАТИЕ ]\n ")

	// Путь и файлы
	path := base.Dir("compress/..")
	src := filepath.Join(path, "go.sum")
	dst := filepath.Join(path, "compress", "go.sum.gz")

	// Исходный файл
	srcFile, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	// Информация исходного файла
	srcInfo, err := srcFile.Stat()
	if err != nil {
		log.Fatal(err)
	}
	srcSize := srcInfo.Size()

	// Чтение исходного файл
	srcBody, err := io.ReadAll(srcFile)
	if err != nil {
		log.Fatal(err)
	}

	// Сжатый файл
	dstFile, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(dst)
	defer dstFile.Close()

	// Сжатие
	w := gzip.NewWriter(dstFile)
	w.Write(srcBody)
	w.Flush()

	// Информация сжатого файла
	dstInfo, err := dstFile.Stat()
	if err != nil {
		log.Fatal(err)
	}
	dstSize := dstInfo.Size()

	// Вывод статистики
	fmt.Printf("Исходный размер: %2.1f Кб\n",
		float64(srcSize)/1000)
	fmt.Printf("Сжатый размер: %2.1f Кб\n",
		float64(dstSize)/1000)
	fmt.Printf("Степень сжатия: %2.2f%%\n",
		float64(dstSize)/float64(srcSize)*100)
}
