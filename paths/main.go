package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(" \n[ ПУТИ ]\n ")

	/* Пути */

	fmt.Println("Объединение:")
	path := filepath.Join("dir", "sub", "file")
	fmt.Println(path)
	fmt.Println(filepath.Join("sub//", "file"))
	fmt.Println(filepath.Join("dir/../dir", "file"))
	fmt.Println()

	fmt.Println("Преобразование:")
	fmt.Println(filepath.ToSlash(path))
	fmt.Println(filepath.Split(path))
	fmt.Println(filepath.Split("dir/sub/file"))
	fmt.Println()

	fmt.Println("Фрагменты:")
	fmt.Println("Dir:", filepath.Dir(path))
	fmt.Println("Base:", filepath.Base(path))
	fmt.Println()

	fmt.Println("Относительный:", filepath.IsAbs("dir/file"))
	fmt.Println("Абсолютный (Linux):", filepath.IsAbs("/dir/file"))
	fmt.Println("Абсолютный (Windows):", filepath.IsAbs("C:\\Go\\Lang"))
	fmt.Println()

	/* Файлы */

	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println("Файл:", filename)
	fmt.Println("Имя:", strings.TrimSuffix(filename, ext))
	fmt.Println("Расширение:", ext)
	fmt.Println()

	/* Отношения */

	rel, err := filepath.Rel("dir/sub1", "dir/sub1/sub2/file")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dir/sub1 > sub2/file:", rel)

	rel, err = filepath.Rel("dir/sub1", "dir/sub2/sub3/file")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dir/sub1 <> sub2/sub3/file:", rel)
}
