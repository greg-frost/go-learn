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
func process(r io.Reader) error {
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
func openFile(filename string) error {
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer r.Close()

	return process(r)
}

// Открытие архива
func openGzipFile(filename string) error {
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

	return process(gz)
}

func main() {
	fmt.Println(" \n[ ЧТЕНИЕ READER ]\n ")

	path := base.Dir("reader/..")

	fmt.Println("Чтение файла:")
	err := openFile(filepath.Join(path, "hello", "main.go"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	fmt.Println("Чтение архива:")
	err = openGzipFile(filepath.Join(path, "reader", "data", "main.tar.gz"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
}
