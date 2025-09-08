package base

import (
	"os"
	"path/filepath"
	"strings"
)

// Получение абсолютного пути
func Dir(target string) string {
	// Текущий каталог
	curr, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Преобразование пути
	curr = filepath.ToSlash(curr)
	target = filepath.ToSlash(target)

	// Разбиение пути
	currDirs := strings.Split(curr, "/")
	targetDirs := strings.Split(target, "/")

	// Сравнение путей
	var p1, p2 int
	n1, n2 := len(currDirs), len(targetDirs)
	for p1 < n1 && p2 < n2 {
		if currDirs[p1] == targetDirs[p2] {
			targetDirs = targetDirs[1:]
			p2++
		}
		p1++
	}

	// Дополнение текущего пути
	if len(targetDirs) > 0 {
		curr = filepath.Join(
			curr,
			filepath.Join(targetDirs...),
		)
	}

	return curr
}
