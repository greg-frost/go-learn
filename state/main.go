package main

import (
	"fmt"
)

// Структура "товар"
type Product struct {
	Title    string
	Price    float32
	Quantity int
	state    State
}

// Конструктор товара
func NewProduct(title string, price float32, quantity int) *Product {
	return &Product{
		Title:    title,
		Price:    price,
		Quantity: quantity,
	}
}

// Статус товара
func (p *Product) State() State {
	return p.state
}

// Интерфейс "состояние"
type State interface {
	Order()
	Pay()
	Deliver()
	Recieve()
	Cancel()
	Return()
}

func main() {
	fmt.Println(" \n[ СОСТОЯНИЕ ]\n ")
}
