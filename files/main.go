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

	"go-learn/base"
)

// Разделитель
var sep = "   "

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

	fmt.Printf("Скопировано %d байт в \"%s\"\n", n, dst)
	return nil
}

// Печать содержимого каталога
func printDir(path string, predicate func(string) bool) {
	var walk func(string)
	walk = func(path string) {
		files, err := os.ReadDir(path)
		check(err)

		for _, f := range files {
			filename := filepath.Join(path, f.Name())
			if predicate(filename) {
				fmt.Println(sep + filename)
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

	path := base.Dir("files/..")
	filename := "hello/main.go"

	// Смена директории
	err := os.Chdir(path)
	check(err)

	// Чтение файла
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	stat, err := file.Stat()
	check(err)

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	check(err)

	str1 := string(bs)
	lines := strings.Split(str1, "\n")
	fmt.Printf("Файл \"%s\":\n", filename)
	for _, line := range lines {
		fmt.Println(sep + line)
	}

	// Второй способ
	bs, err = os.ReadFile(filename)
	check(err)
	str2 := string(bs)

	// Третий способ (устарел)
	bs, err = ioutil.ReadFile(filename)
	check(err)
	str3 := string(bs)

	// Четвертый способ
	_, err = file.Seek(0, io.SeekStart)
	check(err)

	r := bufio.NewReader(file)
	bs, err = r.Peek(int(stat.Size()))
	check(err)
	str4 := string(bs)

	// Сравнение файлов
	if str1 == str2 && str2 == str3 && str3 == str4 {
		fmt.Println("Файлы идентичны!")
	} else {
		fmt.Println("Файлы отличаются...")
	}

	// Создание файла
	filename += ".txt"
	filenameOld := filename + ".old"
	file, err = os.Create(filename)
	check(err)

	file.Write(bs)
	file.WriteString(str2)
	file.Sync()
	fmt.Println("Файл создан и записан!")

	// Второй способ
	os.WriteFile(filename, bs, 0644)
	check(err)

	// Третий способ (дозапись)
	w := bufio.NewWriter(file)
	_, err = w.WriteString("\n" + str2)
	check(err)
	w.Flush()

	// Копирование файла
	copyFile(filename, filenameOld)
	fmt.Println()

	// Очистка
	defer os.Remove(filenameOld)
	defer os.Remove(filename)
	defer file.Close()

	// Чтение каталога
	dir, err := os.Open(path)
	check(err)
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	check(err)

	fmt.Println("Содержимое текущего каталога:")
	for _, fi := range fileInfos[:10] {
		fmt.Printf("%s%s [%s, %d, %t]\n", sep, fi.Name(), fi.ModTime(), fi.Size(), fi.IsDir())
	}
	fmt.Println(sep + "...")
	fmt.Println()

	// Рекурсивный обход
	recPath := "testing"
	fmt.Printf("Все файлы каталога \"%s\":\n", recPath)
	filepath.Walk(recPath, func(path string, info os.FileInfo, err error) error {
		fmt.Println(sep + path)
		return nil
	})
	fmt.Println()

	// Фильтрация
	fmt.Printf("Только тесты каталога \"%s\":\n", recPath)
	printDir(recPath, containsTest)

	// Создание каталога
	fmt.Println()
	dirname := "temp"
	err = os.Mkdir(dirname, 0755)
	check(err)
	fmt.Printf("Создан каталог \"%s\"\n", dirname)

	err = os.MkdirAll(dirname+"/inner", 0755)
	check(err)
	fmt.Printf("Создан каталог \"%s\"\n", dirname+"/inner")

	defer os.RemoveAll(dirname)
	fmt.Println("Созданные каталоги удалены")
	fmt.Println()

	// Временный файл
	tempFile, err := os.CreateTemp("", "temp")
	check(err)
	defer os.Remove(tempFile.Name())
	fmt.Println("Создан временный файл:")
	fmt.Println(sep + tempFile.Name())

	// Временный каталог
	tempDir, err := os.MkdirTemp("", "tempdir")
	check(err)
	defer os.RemoveAll(tempDir)
	fmt.Println("Создан временный каталог:")
	fmt.Println(sep + tempDir)
}
