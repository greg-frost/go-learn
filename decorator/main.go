package main

import (
	"fmt"
)

// Интерфейс "декорируемое"
type Decorable interface {
	Cost() float32
	Description() string
}

// Интерфейс "декоратор"
type Decorator interface {
	Decorable
	Decorate(decorator Decorable)
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
// type Car struct {
// 	cost        float32
// 	description string
// }

// Структура "дом"
// type Home struct {
// 	cost        float32
// 	description string
// }

// Структура "налог"
type Tax struct {
	decorator   Decorable
	cost        float32
	description string
}

// Конструктор "налога"
func NewTax(cost float32, description string) Decorator {
	return &Tax{
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

// Декорирование налога
func (t *Tax) Decorate(decorator Decorable) {
	t.decorator = decorator
}

// Структура "премия"
type Bonus struct {
	decorator   Decorable
	cost        float32
	description string
}

// Конструктор премии
func NewBonus(cost float32, description string) Decorator {
	return &Bonus{
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

// Декорирование премии
func (b *Bonus) Decorate(decorator Decorable) {
	b.decorator = decorator
}

// Структура "гарантия"
// type Warranty struct {
//  decorator   Decorable
// 	cost        float32
// 	description string
// }

// Структура "страховка"
// type Insurance struct {
//  decorator   Decorable
// 	cost        float32
// 	description string
// }

func main() {
	fmt.Println(" \n[ ДЕКОРАТОР ]\n ")

	// Работа, налог, премия
	job := NewJob(250000, "Go-разработчик")
	tax := NewTax(0.13, "Подоходный налог")
	tax.Decorate(job)
	bonus := NewBonus(0.05, "Премия")
	bonus.Decorate(tax)

	fmt.Println("Работа")
	fmt.Printf("Зарплата: %.2f руб.\n", bonus.Cost())
	fmt.Println("Описание:", bonus.Description())
}
