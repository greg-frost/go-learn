package main

import (
	"fmt"
)

// Интерфейс "команда"
type Command interface {
	Execute()
}

// Тип "очередь команд"
type Queue []Command

// Добавление команды
func (q *Queue) Add(command Command) {
	*q = append(*q, command)
}

// Удаление последней команды
func (q *Queue) Remove() {
	if len(*q) > 0 {
		*q = (*q)[:len(*q)-1]
	}
}

// Выполнение всех команд
func (q *Queue) Execute() {
	for _, command := range *q {
		command.Execute()
	}
	*q = nil // Обнуление очереди
}

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

// Структура "команда включения компьютера"
type ComputerOnCommand struct {
	Computer *Computer
}

// Конструктор команды включения компьютера
func NewComputerOnCommand(computer *Computer) *ComputerOnCommand {
	return &ComputerOnCommand{
		Computer: computer,
	}
}

// Выполнение команды включения компьютера
func (c *ComputerOnCommand) Execute() {
	c.Computer.TurnOn()
}

// Структура "команда выключения компьютера"
type ComputerOffCommand struct {
	Computer *Computer
}

// Конструктор команды выключения компьютера
func NewComputerOffCommand(computer *Computer) *ComputerOffCommand {
	return &ComputerOffCommand{
		Computer: computer,
	}
}

// Выполнение команды выключения компьютера
func (c *ComputerOffCommand) Execute() {
	c.Computer.CancelAll()
	c.Computer.TurnOn()
}

func main() {
	fmt.Println(" \n[ КОМАНДА ]\n ")

	computer := NewComputer("GregoryPC")

	var queue Queue
	queue.Add(NewComputerOnCommand(computer))
	queue.Add(NewComputerOffCommand(computer))
	queue.Execute()
}
