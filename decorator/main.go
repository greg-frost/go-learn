package main

import (
	"fmt"
)

// Интерфейс "декорируемое"
type Decorable interface {
	Cost() float32
	Description() string
}

// Структура "работа"
type Job struct {
	cost        float32
	description string
}

// Конструктор работы
func NewJob(cost float32, description string) Decorable {
	return &Job{
		cost:        cost,
		description: description,
	}
}

// Стоимость работы
func (j *Job) Cost() float32 {
	return j.cost
}

// Описание работы
func (j *Job) Description() string {
	return j.description
}

// Структура "машина"
type Car struct {
	cost        float32
	description string
}

// Конструктор машины
func NewCar(cost float32, description string) Decorable {
	return &Car{
		cost:        cost,
		description: description,
	}
}

// Стоимость машины
func (c *Car) Cost() float32 {
	return c.cost
}

// Описание машины
func (c *Car) Description() string {
	return c.description
}

// Структура "дом"
type House struct {
	cost        float32
	description string
}

// Конструктор дома
func NewHouse(cost float32, description string) Decorable {
	return &House{
		cost:        cost,
		description: description,
	}
}

// Стоимость дома
func (h *House) Cost() float32 {
	return h.cost
}

// Описание дома
func (h *House) Description() string {
	return h.description
}

// Структура "налог"
type Tax struct {
	decorator   Decorable
	cost        float32
	description string
}

// Конструктор налога
func NewTax(cost float32, description string, decorator Decorable) Decorable {
	return &Tax{
		decorator:   decorator,
		cost:        cost,
		description: description,
	}
}

// Стоимость налога
func (t *Tax) Cost() float32 {
	return t.decorator.Cost() - t.cost*t.decorator.Cost()
}

// Описание налога
func (t *Tax) Description() string {
	return t.decorator.Description() + " - " + t.description
}

// Структура "премия"
type Bonus struct {
	decorator   Decorable
	cost        float32
	description string
}

// Конструктор премии
func NewBonus(cost float32, description string, decorator Decorable) Decorable {
	return &Bonus{
		decorator:   decorator,
		cost:        cost,
		description: description,
	}
}

// Стоимость премии
func (b *Bonus) Cost() float32 {
	return b.decorator.Cost() + b.cost*b.decorator.Cost()
}

// Описание премии
func (b *Bonus) Description() string {
	return b.decorator.Description() + " + " + b.description
}

// Структура "гарантия"
type Warranty struct {
	decorator   Decorable
	cost        float32
	description string
}

// Конструктор гарантии
func NewWarranty(cost float32, description string, decorator Decorable) Decorable {
	return &Warranty{
		decorator:   decorator,
		cost:        cost,
		description: description,
	}
}

// Стоимость гарантии
func (w *Warranty) Cost() float32 {
	return w.decorator.Cost() - w.cost*w.decorator.Cost()
}

// Описание гарантии
func (w *Warranty) Description() string {
	return w.decorator.Description() + " - " + w.description
}

// Структура "страховка"
type Insurance struct {
	decorator   Decorable
	cost        float32
	description string
}

// Конструктор страховки
func NewInsurance(cost float32, description string, decorator Decorable) Decorable {
	return &Insurance{
		decorator:   decorator,
		cost:        cost,
		description: description,
	}
}

// Стоимость страховки
func (i *Insurance) Cost() float32 {
	return i.decorator.Cost() - i.cost*i.decorator.Cost()
}

// Описание страховки
func (i *Insurance) Description() string {
	return i.decorator.Description() + " - " + i.description
}

func main() {
	fmt.Println(" \n[ ДЕКОРАТОР ]\n ")

	// Работа, налог, премия
	job := NewJob(250000, "Go-разработчик")
	tax := NewTax(0.13, "Подоходный налог", job)
	bonus := NewBonus(0.05, "Премия", tax)

	fmt.Println("Работа")
	fmt.Printf("Зарплата: %.2f руб.\n", bonus.Cost())
	fmt.Println("Описание:", bonus.Description())
}
