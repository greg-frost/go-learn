package main

import (
	"fmt"
)

// Интерфейс "композит"
type Composite interface {
	Add(composite Composite)
	Remove(composite Composite)
}

// Структура "файл"
type File struct {
	Name string
	Size int
}

// Конструктор файла
func NewFile(name string, size int) *File {
	return &File{
		Name: name,
		Size: size,
	}
}

// Добавление композита в файл
func (*File) Add(composite Composite) {}

// Удаление композита из файла
func (*File) Remove(composite Composite) {}

// Структура "каталог"
type Dir struct {
	Name       string
	Composites []Composite
}

// Конструктор каталога
func NewDir(name string) *Dir {
	return &Dir{
		Name: name,
	}
}

// Добавление композита в каталог
func (d *Dir) Add(composite Composite) {
	d.Composites = append(d.Composites, composite)
}

// Удаление композита из каталога
func (d *Dir) Remove(composite Composite) {
	for i, v := range d.Composites {
		if v == composite {
			copy(d.Composites[i:], d.Composites[i+1:])
			d.Composites = d.Composites[:len(d.Composites)-1]
			break
		}
	}
}

func main() {
	fmt.Println(" \n[ КОМПОНОВЩИК ]\n ")

	// Каталоги и файлы
	dir := NewDir("dir")
	sub := NewDir("sub")
	file1 := NewFile("file1.txt", 1242)
	file2 := NewFile("file2.log", 10233)
	file3 := NewFile("file3", 5001)

	// Добавление композитов
	dir.Add(sub)
	dir.Add(file1)
	dir.Add(file2)
	dir.Remove(file2)
	sub.Add(file2)
	sub.Add(file3)

	// Вывод
	fmt.Println("dir:", dir)
	fmt.Println("sub:", sub)
}
