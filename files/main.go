package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Копирование файла
func copyFile(src, dst string) error {
	srcF, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcF.Close()

	dstF, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstF.Close()

	n, err := io.Copy(dstF, srcF)
	if err != nil {
		return err
	}

	fmt.Printf("\nСкопировано %d байт в \"%s\"\n", n, dst)
	return nil
}

// Печать содержимого каталога
func printDir(path string, predicate func(string) bool) {
	var walk func(string)

	walk = func(path string) {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			fmt.Println("ОШИБКА: Невозможно получить содержимое каталога", err)
			return
		}

		for _, f := range files {
			filename := filepath.Join(path, f.Name())

			if predicate(filename) {
				fmt.Println(filename)
			}

			if f.IsDir() {
				walk(filename)
			}
		}
	}

	walk(path)
}

// Поиск "." в строке
func containsDot(s string) bool {
	return strings.Contains(s, ".")
}

// Поиск "_test" в строке
func containsTest(s string) bool {
	return strings.Contains(s, "_test")
}

func main() {
	fmt.Println(" \n[ ФАЙЛЫ ]\n ")

	path := os.Getenv("GOPATH") + "/src/golearn/"
	filename := "hello/main.go"

	/* Первый способ чтения файла */

	file, err := os.Open(path + filename)
	if err != nil {
		fmt.Println("ОШИБКА: Невозможно открыть файл", filename)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println("ОШИБКА: Невозможно получить статистику файла", filename)
		return
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("ОШИБКА: Невозможно прочитать файл", filename)
		return
	}

	str1 := string(bs)
	fmt.Print("Файл ", filename, ":")
	fmt.Println(" \n ")
	fmt.Println(str1)

	/* Второй способ чтения файла */

	bs, err = ioutil.ReadFile(path + filename)
	if err != nil {
		fmt.Println("ОШИБКА: Невозможно открыть файл", filename)
		return
	}

	str2 := string(bs)
	if str1 == str2 {
		fmt.Println("Файлы идентичны!")
	} else {
		fmt.Println("Файлы отличаются...")
	}

	/* Создание файла */

	filename += ".txt"
	filenameOld := filename + ".old"
	file, err = os.Create(path + filename)
	if err != nil {
		fmt.Println("ОШИБКА: Невозможно создать файл", filename)
		return
	}
	defer os.Remove(path + filenameOld)
	defer os.Remove(path + filename)
	defer copyFile(path+filename, path+filenameOld)
	defer file.Close()

	file.WriteString(str2)
	fmt.Println("Файл создан и записан!")

	fmt.Println()

	/* Чтение каталога */

	dir, err := os.Open(path)
	if err != nil {
		fmt.Println("ОШИБКА: Невозможно открыть текущий каталог")
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("ОШИБКА: Невозможно прочитать текущий каталог")
		return
	}

	fmt.Println("Содержимое текущего каталога:")
	fmt.Println()
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	fmt.Println()

	/* Рекурсивный обход каталога */

	recPath := "C:/Go/Lang/src/unicode/"

	fmt.Printf("Полное содержимое каталога \"%s\":\n\n", recPath)
	filepath.Walk(recPath, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

	fmt.Println()

	fmt.Printf("Отфильтованное содержимое каталога \"%s\":\n\n", recPath)
	printDir(recPath, containsTest)
}
