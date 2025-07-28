package main

import (
	"fmt"
)

// Структура "компьютер"
type Computer struct {
	Name string
}

// Конструктор компьютера
func NewComputer(name string) *Computer {
	return &Computer{
		Name: name,
	}
}

// Включение компьютера
func (c *Computer) TurnOn() {
	fmt.Printf("Компьютер %s: включение\n", c.Name)
}

// Отмена всех операций компьютера
func (c *Computer) CancelAll() {
	fmt.Printf("Компьютер %s: отмена всех операций\n", c.Name)
}

// Выключение компьютера
func (c *Computer) TurnOff() {
	fmt.Printf("Компьютер %s: выключение\n", c.Name)
}

// Структура "плеер"
type Player struct {
	Name   string
	Disk   string
	Volume int
}

// Конструктор плеера
func NewPlayer(name string) *Player {
	return &Player{
		Name: name,
	}
}

// Включение плеера
func (c *Player) TurnOn() {
	fmt.Printf("Плеер %s: включение\n", c.Name)
}

// Загрузка диска в плеер
func (c *Player) InsertDisk(disk string) {
	c.Disk = disk
	fmt.Printf("Плеер %s: загрузка диска %q\n", c.Name, c.Disk)
}

// Извлечение диска из плеера
func (c *Player) EjectDisk() {
	fmt.Printf("Плеер %s: извлечение диска %q\n", c.Name, c.Disk)
	c.Disk = ""
}

// Установка громкости плеера
func (c *Player) SetVolume(volume int) {
	c.Volume = volume
	fmt.Printf("Плеер %s: установка громкости %d\n", c.Name, c.Volume)
}

// Проигрывание трека в плеере
func (c *Player) Play() {
	fmt.Printf("Плеер %s: проигрывание диска %q\n", c.Name, c.Disk)
}

// Пауза трека в плеере
func (c *Player) Pause() {
	fmt.Printf("Плеер %s: пауза диска %q\n", c.Name, c.Disk)
}

// Остановка трека в плеере
func (c *Player) Stop() {
	fmt.Printf("Плеер %s: остановка диска %q\n", c.Name, c.Disk)
}

// Очистка очереди проигрывания плеера
func (c *Player) Clear() {
	fmt.Printf("Плеер %s: очистка очереди проигрывания\n", c.Name)
}

// Включение плеера
func (c *Player) TurnOff() {
	fmt.Printf("Плеер %s: включение\n", c.Name)
}

func main() {
	fmt.Println(" \n[ ФАСАД ]\n ")

	// Компьютер и плеер
	computer := NewComputer("GregoryPC")
	player := NewPlayer("Vinyl")

	// Включение музыки
	computer.TurnOn()
	player.SetVolume(10)
	player.InsertDisk("Нить Ариадны")
	player.Play()

	// Выключение музыки
	player.Pause()
	player.Stop()
	player.EjectDisk()
	player.SetVolume(0)
	computer.CancelAll()
	computer.TurnOff()
}
