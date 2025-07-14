package main

import (
	"fmt"
	"strings"
)

// Интерфейс "ОС"
type OS interface {
	Run()
	Stop()
}

// Структура "Windows"
type Windows struct {
	Name    string
	Version int
}

// Конструктор Windows
func NewWindows(name string, version int) OS {
	return &Windows{
		Name:    name,
		Version: version,
	}
}

// Запуск Windows
func (w *Windows) Run() {
	fmt.Printf("%s %d: запуск...\n", w.Name, w.Version)
}

// Остановка Windows
func (w *Windows) Stop() {
	fmt.Printf("%s %d: остановка...\n", w.Name, w.Version)
}

// Структура "Linux"
type Linux struct {
	Name    string
	Version float32
}

// Конструктор Linux
func NewLinux(name string, version float32) OS {
	return &Linux{
		Name:    name,
		Version: version,
	}
}

// Запуск Linux
func (l *Linux) Run() {
	fmt.Printf("%s %.2f: запуск...\n", l.Name, l.Version)
}

// Остановка Linux
func (l *Linux) Stop() {
	fmt.Printf("%s %.2f: остановка...\n", l.Name, l.Version)
}

// Структура "MacOS"
type MacOS struct {
	Name    string
	Version float32
}

// Конструктор MacOS
func NewMacOS(name string, version float32) OS {
	return &MacOS{
		Name:    name,
		Version: version,
	}
}

// Запуск MacOS
func (m *MacOS) Run() {
	fmt.Printf("%s %.1f: запуск...\n", m.Name, m.Version)
}

// Остановка MacOS
func (m *MacOS) Stop() {
	fmt.Printf("%s %.1f: остановка...\n", m.Name, m.Version)
}

// Простая фабрика
func SimpleFactory(os string) OS {
	os = strings.ToLower(os)
	os = strings.ReplaceAll(os, " ", "")
	switch os {
	case "windows":
		return NewWindows("Windows", 10)
	case "linux":
		return NewLinux("Ubuntu", 24.04)
	case "macos":
		return NewMacOS("Mac OS", 15.5)
	default:
		return NewWindows("Windows", 7)
	}
}

func main() {
	fmt.Println(" \n[ ФАБРИКА ]\n ")

	// Простая фабрика
	fmt.Println("Простая фабрика:")
	fmt.Println()
	os := SimpleFactory("Windows")
	os.Run()
	fmt.Println("Выполнение операций...")
	os.Stop()
}
