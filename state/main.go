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

// Изменить статус товара
func (p *Product) ChangeState(state State) {
	p.state = state
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

// Структура "состояние - товар готов к покупке"
type ReadyState struct {
	product *Product
}

// Конструктор состояния
func NewReadyState(product *Product) *ReadyState {
	return &ReadyState{
		product: product,
	}
}

// Заказ
func (r *ReadyState) Order() {
	if r.product.Quantity > 0 {
		fmt.Printf("Товар %s успешно заказан\n", r.product.Title)
		r.product.Quantity--
		r.product.ChangeState(NewOrderedState(r.product))
	} else {
		fmt.Printf("Товар %s закончился\n", r.product.Title)
		// r.product.ChangeState(NewOutOfStockState(r.product))
	}
}

// Оплата
func (r *ReadyState) Pay() {
	fmt.Println("Невозможно оплатить: товар еще не заказан")
}

// Доставка
func (r *ReadyState) Deliver() {
	fmt.Println("Невозможно доставить: товар еще не заказан")
}

// Получение
func (r *ReadyState) Recieve() {
	fmt.Println("Невозможно получить: товар еще не заказан")
}

// Отмена
func (r *ReadyState) Cancel() {
	fmt.Println("Невозможно отменить: товар еще не заказан")
}

// Возврат
func (r *ReadyState) Return() {
	fmt.Println("Невозможно вернуть: товар еще не заказан")
}

// Структура "состояние - товар заказан"
type OrderedState struct {
	product *Product
}

// Конструктор состояния
func NewOrderedState(product *Product) *OrderedState {
	return &OrderedState{
		product: product,
	}
}

// Заказ
func (o *OrderedState) Order() {
	fmt.Println("Невозможно заказать: товар уже заказан")
}

// Оплата
func (o *OrderedState) Pay() {
	fmt.Printf("Товар %s успешно оплачен: %.2f руб.\n", o.product.Title, o.product.Price)
	// o.product.ChangeState(NewOrderedState(o.product))
}

// Доставка
func (o *OrderedState) Deliver() {
	fmt.Println("Невозможно доставить: товар еще не оплачен")
}

// Получение
func (o *OrderedState) Recieve() {
	fmt.Println("Невозможно получить: товар еще не оплачен")
}

// Отмена
func (o *OrderedState) Cancel() {
	fmt.Printf("Заказ товара %s отменен, деньги возвращены: %.2f руб.\n",
		o.product.Title, o.product.Price)
	o.product.Quantity++
	o.product.ChangeState(NewReadyState(o.product))
}

// Возврат
func (o *OrderedState) Return() {
	fmt.Println("Невозможно вернуть: товар еще не оплачен")
}

func main() {
	fmt.Println(" \n[ СОСТОЯНИЕ ]\n ")
}
