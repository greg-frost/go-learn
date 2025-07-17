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
func NewComputerOnCommand(computer *Computer) Command {
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
func NewComputerOffCommand(computer *Computer) Command {
	return &ComputerOffCommand{
		Computer: computer,
	}
}

// Выполнение команды выключения компьютера
func (c *ComputerOffCommand) Execute() {
	c.Computer.CancelAll()
	c.Computer.TurnOff()
}

// Структура "принтер"
type Printer struct {
	Name string
}

// Конструктор принтера
func NewPrinter(name string) *Printer {
	return &Printer{
		Name: name,
	}
}

// Включение принтера
func (c *Printer) TurnOn() {
	fmt.Printf("Принтер %s: включение\n", c.Name)
}

// Разогрев принтера
func (c *Printer) WarmUp() {
	fmt.Printf("Принтер %s: разогрев барабана\n", c.Name)
}

// Загрузка бумаги в принтер
func (c *Printer) LoadPaper() {
	fmt.Printf("Принтер %s: загрузка бумаги\n", c.Name)
}

// Печать в принтере
func (c *Printer) Print() {
	fmt.Printf("Принтер %s: печать\n", c.Name)
}

// Очистка очереди печати принтера
func (c *Printer) ClearQueue() {
	fmt.Printf("Принтер %s: очистка очереди печати\n", c.Name)
}

// Выключение принтера
func (c *Printer) TurnDown() {
	fmt.Printf("Принтер %s: выключение\n", c.Name)
}

// Структура "команда включения принтера"
type PrinterOnCommand struct {
	Printer *Printer
}

// Конструктор команды включения принтера
func NewPrinterOnCommand(printer *Printer) Command {
	return &PrinterOnCommand{
		Printer: printer,
	}
}

// Выполнение команды включения принтера
func (c *PrinterOnCommand) Execute() {
	c.Printer.TurnOn()
	c.Printer.WarmUp()
	c.Printer.LoadPaper()
}

// Структура "команда печати для принтера"
type PrinterPrintCommand struct {
	Printer *Printer
}

// Конструктор команды печати для принтера
func NewPrinterPrintCommand(printer *Printer) Command {
	return &PrinterPrintCommand{
		Printer: printer,
	}
}

// Выполнение команды печати для принтера
func (c *PrinterPrintCommand) Execute() {
	c.Printer.Print()
}

// Структура "команда выключения принтера"
type PrinterOffCommand struct {
	Printer *Printer
}

// Конструктор команды выключения принтера
func NewPrinterOffCommand(printer *Printer) Command {
	return &PrinterOffCommand{
		Printer: printer,
	}
}

// Выполнение команды выключения принтера
func (c *PrinterOffCommand) Execute() {
	c.Printer.ClearQueue()
	c.Printer.TurnDown()
}

func main() {
	fmt.Println(" \n[ КОМАНДА ]\n ")

	// Компьютер и принтер
	computer := NewComputer("GregoryPC")
	printer := NewPrinter("Xerox Phaser 3117")

	// Очередь команд
	queue := NewQueue()
	queue.Add(NewComputerOnCommand(computer))
	queue.Add(NewPrinterOnCommand(printer))
	queue.Add(NewPrinterPrintCommand(printer))
	queue.Add(NewPrinterOffCommand(printer))
	queue.Add(NewComputerOffCommand(computer))
	queue.Execute()
}
