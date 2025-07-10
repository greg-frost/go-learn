package main

import (
	"fmt"
)

// Интерфейс "строитель"
type Builder interface {
	CPU(frequency float32) Builder
	RAM(gigabytes int) Builder
	SSD(volume int) Builder

	Build() Computer
}

// Структура "компьютер"
type Computer struct {
	CPU float32
	RAM int
	SSD int
}

// Структура "строитель компьютера"
type computerBuilder struct {
	cpu float32
	ram int
	ssd int
}

// Конструктор
func NewComputerBuilder() Builder {
	return computerBuilder{}
}

// Установка CPU
func (cb computerBuilder) CPU(frequency float32) Builder {
	cb.cpu = frequency
	return cb
}

// Установка RAM
func (cb computerBuilder) RAM(gigabytes int) Builder {
	cb.ram = gigabytes
	return cb
}

// Установка SSD
func (cb computerBuilder) SSD(volume int) Builder {
	cb.ssd = volume
	return cb
}

// Построение
func (cb computerBuilder) Build() Computer {
	if cb.cpu == 0 {
		cb.cpu = 4.0
	}
	if cb.ram == 0 {
		cb.ram = 24
	}
	if cb.ssd == 0 {
		cb.ssd = 2000
	}
	return Computer{
		CPU: cb.cpu,
		RAM: cb.ram,
		SSD: cb.ssd,
	}
}

func main() {
	fmt.Println(" \n[ СТРОИТЕЛЬ ]\n ")

	// Строитель
	compBuilder := NewComputerBuilder()

	// Построение первого компьютера
	computer1 := compBuilder.CPU(3.2).RAM(8).SSD(500).Build()
	fmt.Println("Компьютер 1")
	fmt.Printf("CPU: %.1f GHz, RAM: %d GB, SSD: %d GB\n",
		computer1.CPU, computer1.RAM, computer1.SSD)

	// Построение второго компьютера
	computer2 := compBuilder.RAM(16).SSD(1000).Build()
	fmt.Println("Компьютер 2")
	fmt.Printf("CPU: %.1f GHz, RAM: %d GB, SSD: %d GB\n",
		computer2.CPU, computer2.RAM, computer2.SSD)

	// Построение третьего компьютера
	computer3 := compBuilder.CPU(4.5).Build()
	fmt.Println("Компьютер 3")
	fmt.Printf("CPU: %.1f GHz, RAM: %d GB, SSD: %d GB\n",
		computer3.CPU, computer3.RAM, computer3.SSD)
}
