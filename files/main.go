package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Проверка
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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
		check(err)

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
	fmt.Println(" \n[ ФАЙЛЫ И ПАПКИ ]\n ")

	path := os.Getenv("GOPATH") + "/src/golearn/"
	filename := "hello/main.go"

	/* Чтение файла */

	// Первый способ
	file, err := os.Open(path + filename)
	check(err)
	defer file.Close()

	stat, err := file.Stat()
	check(err)

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	check(err)

	str1 := string(bs)
	fmt.Printf("Файл %s:\n\n", filename)
	fmt.Printf("\"\"\"\n%s\"\"\"\n\n", str1)

	// Второй способ
	bs, err = ioutil.ReadFile(path + filename)
	check(err)
	str2 := string(bs)

	// Третий способ
	bs, err = os.ReadFile(path + filename)
	check(err)
	str3 := string(bs)

	// Четвертый способ
	_, err = file.Seek(0, io.SeekStart)
	check(err)

	r := bufio.NewReader(file)
	bs, err = r.Peek(int(stat.Size()))
	check(err)
	str4 := string(bs)

	/* Сравнение файлов */

	if str1 == str2 && str2 == str3 && str3 == str4 {
		fmt.Println("Файлы идентичны!")
	} else {
		fmt.Println("Файлы отличаются...")
	}

	/* Создание файла */

	// Первый способ
	filename += ".txt"
	filenameOld := filename + ".old"
	file, err = os.Create(path + filename)
	check(err)

	file.Write(bs)
	file.WriteString(str2)
	fmt.Println("Файл создан и записан!")
	fmt.Println()

	file.Sync()

	// Второй способ
	os.WriteFile(path+filename, bs, 0644)
	check(err)

	// Третий способ (дозапись)
	w := bufio.NewWriter(file)
	_, err = w.WriteString("\n" + str2)
	check(err)

	w.Flush()

	// Очистка
	defer os.Remove(path + filenameOld)
	defer os.Remove(path + filename)
	defer copyFile(path+filename, path+filenameOld)
	defer file.Close()

	/* Чтение каталога */

	dir, err := os.Open(path)
	check(err)
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	check(err)

	fmt.Println("Содержимое текущего каталога:")
	fmt.Println()
	for _, fi := range fileInfos[:10] {
		fmt.Printf("%s [%s, %d, %t]\n", fi.Name(), fi.ModTime(), fi.Size(), fi.IsDir())
	}
	fmt.Println("...")
	fmt.Println()

	// Рекурсивный обход
	recPath := path + "test"
	fmt.Printf("Все файлы каталога \"%s\":\n\n", recPath)
	filepath.Walk(recPath, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
	fmt.Println()

	// Фильтрация
	fmt.Printf("Только тесты каталога \"%s\":\n\n", recPath)
	printDir(recPath, containsTest)
}
