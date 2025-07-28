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

// Выключение плеера
func (c *Player) TurnOff() {
	fmt.Printf("Плеер %s: включение\n", c.Name)
}

type Facade struct {
	Computer *Computer
	Player   *Player
}

func NewFacade(computer *Computer, player *Player) *Facade {
	return &Facade{
		Computer: computer,
		Player:   player,
	}
}

// Проиграть первый альбом
func (f *Facade) PlayFirstAlbum() {
	f.Computer.TurnOn()
	f.Player.SetVolume(8)
	f.Player.InsertDisk("Нить Ариадны")
	f.Player.Play()
}

// Остановить первый альбом
func (f *Facade) StopFirstAlbum() {
	f.Player.Pause()
	f.Player.Stop()
	f.Player.EjectDisk()
	f.Player.SetVolume(0)
	f.Computer.CancelAll()
	f.Computer.TurnOff()
}

// Проиграть второй альбом
func (f *Facade) PlaySecondAlbum() {
	f.Computer.TurnOn()
	f.Player.SetVolume(9)
	f.Player.InsertDisk("Cold Face, Your Grace")
	f.Player.Play()
}

// Остановить второй альбом
func (f *Facade) StopSecondAlbum() {
	f.Player.Pause()
	f.Player.Stop()
	f.Player.EjectDisk()
	f.Player.SetVolume(0)
	f.Computer.CancelAll()
	f.Computer.TurnOff()
}

func main() {
	fmt.Println(" \n[ ФАСАД ]\n ")

	// Компьютер, плеер и фасад
	computer := NewComputer("GregoryPC")
	player := NewPlayer("Vinyl")
	facade := NewFacade(computer, player)

	// Первый альбом
	facade.PlayFirstAlbum()
	facade.StopFirstAlbum()
	fmt.Println()

	// Другие компьютер, плеер и фасад
	computer = NewComputer("Notebook")
	player = NewPlayer("CD")
	facade = NewFacade(computer, player)

	// Второй альбом
	facade.PlaySecondAlbum()
	facade.StopSecondAlbum()
}
