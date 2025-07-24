package main

import (
	"fmt"
)

// Интерфейс "операционная система"
type OS interface {
	Factory(version float32)
	Run()
	Stop()
}

// Структура "Windows"
type Windows struct {
	Name    string
	Version float32
}

// Конструктор Windows
func NewWindows() OS {
	return &Windows{
		Name: "Windows",
	}
}

// Создание Windows
func (w *Windows) Factory(version float32) {
	w.Version = version
	fmt.Printf("%s %.0f: создание\n", w.Name, w.Version)
}

// Запуск Windows
func (w *Windows) Run() {
	fmt.Printf("%s %.0f: запуск\n", w.Name, w.Version)
}

// Остановка Windows
func (w *Windows) Stop() {
	fmt.Printf("%s %.0f: остановка\n", w.Name, w.Version)
}

// Структура "Linux"
type Linux struct {
	Name    string
	Version float32
}

// Конструктор Linux
func NewLinux() OS {
	return &Linux{
		Name: "Linux",
	}
}

// Создание Linux
func (l *Linux) Factory(version float32) {
	l.Version = version
	fmt.Printf("%s %.2f: создание\n", l.Name, l.Version)
}

// Запуск Linux
func (l *Linux) Run() {
	fmt.Printf("%s %.2f: запуск\n", l.Name, l.Version)
}

// Остановка Linux
func (l *Linux) Stop() {
	fmt.Printf("%s %.2f: остановка\n", l.Name, l.Version)
}

// Структура "MacOS"
type MacOS struct {
	Name    string
	Version float32
}

// Конструктор MacOS
func NewMacOS() OS {
	return &MacOS{
		Name: "Mac OS",
	}
}

// Создание MacOS
func (m *MacOS) Factory(version float32) {
	m.Version = version
	fmt.Printf("%s %.1f: создание\n", m.Name, m.Version)
}

// Запуск MacOS
func (m *MacOS) Run() {
	fmt.Printf("%s %.1f: запуск\n", m.Name, m.Version)
}

// Остановка MacOS
func (m *MacOS) Stop() {
	fmt.Printf("%s %.1f: остановка\n", m.Name, m.Version)
}

// Структура "сервер"
type server struct {
	os OS
}

// Конструктор сервера
func NewServer(os OS) *server {
	return &server{
		os: os,
	}
}

// Запуск сервера
func (s *server) Run(version float32) {
	s.os.Factory(version)
	s.os.Run()
}

// Остановка сервера
func (s *server) Stop() {
	s.os.Stop()
}

func main() {
	fmt.Println(" \n[ ФАБРИЧНЫЙ МЕТОД ]\n ")

	server := NewServer(NewLinux())
	server.Run(24.04)
	fmt.Println("Выполнение операций")
	server.Stop()
}
