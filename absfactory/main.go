package main

import (
	"fmt"
)

// Интерфейс "операционная система"
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
	fmt.Printf("%s %d: запуск\n", w.Name, w.Version)
}

// Остановка Windows
func (w *Windows) Stop() {
	fmt.Printf("%s %d: остановка\n", w.Name, w.Version)
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
func NewMacOS(name string, version float32) OS {
	return &MacOS{
		Name:    name,
		Version: version,
	}
}

// Запуск MacOS
func (m *MacOS) Run() {
	fmt.Printf("%s %.1f: запуск\n", m.Name, m.Version)
}

// Остановка MacOS
func (m *MacOS) Stop() {
	fmt.Printf("%s %.1f: остановка\n", m.Name, m.Version)
}

// Интерфейс "язык программирования"
type Language interface {
	Run()
	Stop()
}

// Структура "Go"
type Go struct {
	Name    string
	Version float32
}

// Конструктор Go
func NewGo(name string, version float32) OS {
	return &Go{
		Name:    name,
		Version: version,
	}
}

// Запуск Go
func (g *Go) Run() {
	fmt.Printf("%s %.2f: запуск\n", g.Name, g.Version)
}

// Остановка Go
func (g *Go) Stop() {
	fmt.Printf("%s %.2f: остановка\n", g.Name, g.Version)
}

// Структура "PHP"
type PHP struct {
	Name    string
	Version float32
}

// Конструктор PHP
func NewPHP(name string, version float32) OS {
	return &PHP{
		Name:    name,
		Version: version,
	}
}

// Запуск PHP
func (p *PHP) Run() {
	fmt.Printf("%s %.1f: запуск\n", p.Name, p.Version)
}

// Остановка PHP
func (p *PHP) Stop() {
	fmt.Printf("%s %.1f: остановка\n", p.Name, p.Version)
}

// Интерфейс "база данных"
type Database interface {
	Run()
	Stop()
}

// Структура "PostrgeSQL"
type PostrgeSQL struct {
	Name    string
	Version int
}

// Конструктор PostrgeSQL
func NewPostrgeSQL(name string, version int) OS {
	return &PostrgeSQL{
		Name:    name,
		Version: version,
	}
}

// Запуск PostrgeSQL
func (p *PostrgeSQL) Run() {
	fmt.Printf("%s %d: запуск\n", p.Name, p.Version)
}

// Остановка PostrgeSQL
func (p *PostrgeSQL) Stop() {
	fmt.Printf("%s %d: остановка\n", p.Name, p.Version)
}

// Структура "MySQL"
type MySQL struct {
	Name    string
	Version float32
}

// Конструктор MySQL
func NewMySQL(name string, version float32) OS {
	return &MySQL{
		Name:    name,
		Version: version,
	}
}

// Запуск MySQL
func (m *MySQL) Run() {
	fmt.Printf("%s %.1f: запуск\n", m.Name, m.Version)
}

// Остановка MySQL
func (m *MySQL) Stop() {
	fmt.Printf("%s %.1f: остановка\n", m.Name, m.Version)
}

// Интерфейс "сервер"
type Server interface {
	Run()
	Stop()
}

// Структура "сервер"
type server struct {
	os       OS
	language Language
	database Database
}

// Конструктор сервера
func NewServer(factory Factory) Server {
	return &server{
		os:       factory.CreateOS(),
		language: factory.CreateLanguage(),
		database: factory.CreateDatabase(),
	}
}

// Запуск сервера
func (s *server) Run() {
	s.os.Run()
	s.language.Run()
	s.database.Run()
}

// Остановка сервера
func (s *server) Stop() {
	s.database.Stop()
	s.language.Stop()
	s.os.Stop()
}

// Интерфейс "фабрика"
type Factory interface {
	CreateOS() OS
	CreateLanguage() Language
	CreateDatabase() Database
}

// Структура "фабрика Go"
type GoFactory struct{}

// Конструктор фабрики Go
func NewGoFactory() Factory {
	return &GoFactory{}
}

// Создание операционной системы
func (*GoFactory) CreateOS() OS {
	return NewLinux("Linux", 24.04)
}

// Создание языка программирования
func (*GoFactory) CreateLanguage() Language {
	return NewGo("Go", 1.20)
}

// Создание базы данных
func (*GoFactory) CreateDatabase() Database {
	return NewPostrgeSQL("PostgreSQL", 17)
}

// Структура "фабрика PHP"
type PhpFactory struct{}

// Конструктор фабрики PHP
func NewPhpFactory() Factory {
	return &PhpFactory{}
}

// Создание операционной системы
func (*PhpFactory) CreateOS() OS {
	return NewMacOS("Mac OS", 15.5)
}

// Создание языка программирования
func (*PhpFactory) CreateLanguage() Language {
	return NewPHP("PHP", 8.4)
}

// Создание базы данных
func (*PhpFactory) CreateDatabase() Database {
	return NewMySQL("MySQL", 9.3)
}

func main() {
	fmt.Println(" \n[ АБСТРАКТНАЯ ФАБРИКА ]\n ")

	factory := NewPhpFactory()
	server := NewServer(factory)
	server.Run()
	fmt.Println("Выполнение операций")
	server.Stop()
}
