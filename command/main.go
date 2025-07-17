package main

import (
	"fmt"
)

// Интерфейс "команда"
type Command interface {
	Execute()
}

// Структура "очередь команд"
type Queue struct {
	Commands []Command
}

// Конструктор очереди команд
func NewQueue() *Queue {
	return new(Queue)
}

// Добавление команды
func (q *Queue) Add(command Command) {
	q.Commands = append(q.Commands, command)
}

// Удаление последней команды
func (q *Queue) RemoveLast() {
	if len(q.Commands) > 0 {
		q.Commands = q.Commands[:len(q.Commands)-1]
	}
}

// Очистка очереди
func (q *Queue) ClearAll() {
	q.Commands = nil
}

// Выполнение всех команд
func (q *Queue) Execute() {
	for _, command := range q.Commands {
		command.Execute()
	}
	q.ClearAll()
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

	queue := NewQueue()
	computer := NewComputer("GregoryPC")

	queue.Add(NewComputerOnCommand(computer))
	queue.Add(NewComputerOffCommand(computer))
	queue.Execute()
}
