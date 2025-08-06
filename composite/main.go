package main

import (
	"fmt"
)

// Интерфейс "композит"
type Composite interface {
	Add(composite Composite)
	Remove(composite Composite)
	IsDir() bool
	Size() int
}

// Структура "файл"
type File struct {
	name string
	size int
}

// Конструктор файла
func NewFile(name string, size int) *File {
	return &File{
		name: name,
		size: size,
	}
}

// Добавление композита в файл
func (*File) Add(composite Composite) {}

// Удаление композита из файла
func (*File) Remove(composite Composite) {}

// Каталог ли файл
func (*File) IsDir() bool {
	return false
}

// Размер файла
func (f *File) Size() int {
	return f.size
}

// Структура "каталог"
type Dir struct {
	name       string
	composites []Composite
}

// Конструктор каталога
func NewDir(name string) *Dir {
	return &Dir{
		name: name,
	}
}

// Добавление композита в каталог
func (d *Dir) Add(composite Composite) {
	d.composites = append(d.composites, composite)
}

// Удаление композита из каталога
func (d *Dir) Remove(composite Composite) {
	for i, c := range d.composites {
		if c == composite {
			copy(d.composites[i:], d.composites[i+1:])
			d.composites = d.composites[:len(d.composites)-1]
			break
		}
	}
}

// Каталог ли каталог
func (*Dir) IsDir() bool {
	return true
}

// Размер каталога
func (d *Dir) Size() int {
	var size int
	for _, c := range d.composites {
		size += c.Size()
	}
	return size
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

	// Печать размеров
	fmt.Println("Размеры")
	fmt.Println("dir:", dir.Size())
	fmt.Println("sub:", sub.Size())
}
