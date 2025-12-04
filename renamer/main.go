package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

var (
	ext    = ".jpg" // Расширение файла
	format = "%d"   // Формат файла
	order  = "asc"  // Порядок сортировки
	tries  = 3      // Число попыток
)

// Структура "файл"
type File struct {
	name     string
	modified int64
}

func main() {
	fmt.Println(" \n[ ПЕРЕИМЕНОВАНИЕ ]\n ")

	var rename []File

	// Чтение каталога
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	// Обход файлов
	for _, f := range files {
		name := f.Name()

		// Пропуск директорий
		if f.IsDir() {
			continue
		}

		// Пропуск других расширений
		if !strings.HasSuffix(
			strings.ToLower(name),
			strings.ToLower(ext),
		) {
			continue
		}

		// Информация файла
		info, err := f.Info()
		if err != nil {
			continue
		}

		rename = append(rename, File{
			name:     name,
			modified: info.ModTime().Unix(),
		})
	}

	if len(rename) == 0 {
		fmt.Println("Нет файлов для переименования...")
		return
	}

	// Сортировка файлов
	sort.Slice(rename, func(i, j int) bool {
		if strings.ToLower(order) == "desc" {
			return rename[i].modified > rename[j].modified
		}
		return rename[i].modified < rename[j].modified
	})

	// Переименование
	var count int
	for i, f := range rename {
		oldname := f.name
		newname := fmt.Sprintf(format+ext, i+1)
		fmt.Printf("%s -> %s\n", oldname, newname)

		try := tries
		err := os.Rename(oldname, newname)
		for err != nil && try > 0 {
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("%s -x %s (%d)\n", oldname, newname, try)
			err = os.Rename(oldname, newname)
			try--
		}
		count++
	}
	fmt.Println()
	fmt.Println("Переименовано файлов:", count)
}
