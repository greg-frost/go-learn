package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"go-learn/base"
)

// Чтение io.Reader
func Process(r io.Reader) error {
	data := make([]byte, 100)
	for {
		count, err := r.Read(data)
		if err != nil {
			if err != io.EOF {
				return err
			}
			return nil
		}
		if count == 0 {
			return nil
		}
		fmt.Print(string(data[:count]))
	}
}

// Открытие файла
func OpenFile(filename string) error {
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer r.Close()

	return Process(r)
}

// Открытие архива
func OpenGzipFile(filename string) error {
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer r.Close()

	gz, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gz.Close()

	return Process(gz)
}

func main() {
	fmt.Println(" \n[ ЧТЕНИЕ READER ]\n ")

	path := base.Dir("reader/..")

	// Файл
	fmt.Println("Чтение файла:")
	err := OpenFile(filepath.Join(path, "hello", "main.go"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	// Архив
	fmt.Println("Чтение архива:")
	err = OpenGzipFile(filepath.Join(path, "reader", "data", "main.tar.gz"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
}
