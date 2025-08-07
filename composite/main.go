package main

import (
	"fmt"
	"strings"
)

// Интерфейс "композит"
type Composite interface {
	Add(composite Composite)
	Remove(composite Composite)
	Name() string
	Size() int
	IsDir() bool
	Composites() []Composite
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

// Имя файла
func (f *File) Name() string {
	return f.name
}

// Размер файла
func (f *File) Size() int {
	return f.size
}

// Каталог ли файл
func (*File) IsDir() bool {
	return false
}

// Список композитов файла
func (*File) Composites() []Composite {
	return nil
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

// Имя каталога
func (d *Dir) Name() string {
	return d.name
}

// Размер каталога
func (d *Dir) Size() int {
	var size int
	for _, c := range d.composites {
		size += c.Size()
	}
	return size
}

// Каталог ли каталог
func (*Dir) IsDir() bool {
	return true
}

// Список композитов каталога
func (d *Dir) Composites() []Composite {
	return d.composites
}

// Печать композита
func Print(composite Composite) {
	walk(composite, 0)
}

// Обход композитов
func walk(composite Composite, level int) {
	fmt.Print(strings.Repeat("   ", level))
	fmt.Println(composite.Name())
	if composite.IsDir() {
		for _, c := range composite.Composites() {
			walk(c, level+1)
		}
	}
}

func main() {
	fmt.Println(" \n[ КОМПОНОВЩИК ]\n ")

	// Каталоги и файлы
	dir := NewDir("dir")
	sub := NewDir("sub")
	file1 := NewFile("file1.txt", 1024)
	file2 := NewFile("file2.log", 10502)
	file3 := NewFile("file3", 5048)

	// Вложение
	dir.Add(sub)
	dir.Add(file1)
	dir.Add(file2)
	dir.Remove(file2)
	sub.Add(file2)
	sub.Add(file3)

	// Вывод
	fmt.Println("Структура:")
	Print(dir)
	fmt.Println()
	fmt.Println("Размеры:")
	fmt.Println("dir -", dir.Size())
	fmt.Println("sub -", sub.Size())
}
