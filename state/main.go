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

// Структура "товар готов к покупке"
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
func (s *ReadyState) Order() {
	if s.product.Quantity > 0 {
		fmt.Printf("Товар %s успешно заказан\n", s.product.Title)
		s.product.Quantity--
		s.product.ChangeState(NewOrderedState(s.product))
	} else {
		fmt.Printf("Товар %s закончился\n", s.product.Title)
		// r.product.ChangeState(NewOutOfStockState(r.product))
	}
}

// Оплата
func (*ReadyState) Pay() {
	fmt.Println("Невозможно оплатить: товар еще не заказан")
}

// Доставка
func (*ReadyState) Deliver() {
	fmt.Println("Невозможно доставить: товар еще не заказан")
}

// Получение
func (*ReadyState) Recieve() {
	fmt.Println("Невозможно получить: товар еще не заказан")
}

// Отмена
func (*ReadyState) Cancel() {
	fmt.Println("Невозможно отменить: товар еще не заказан")
}

// Возврат
func (*ReadyState) Return() {
	fmt.Println("Невозможно вернуть: товар еще не заказан")
}

// Структура "товар заказан"
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
func (*OrderedState) Order() {
	fmt.Println("Невозможно заказать: товар уже заказан")
}

// Оплата
func (s *OrderedState) Pay() {
	fmt.Printf("Товар %s успешно оплачен: %.2f руб.\n", s.product.Title, s.product.Price)
	s.product.ChangeState(NewPayedState(s.product))
}

// Доставка
func (*OrderedState) Deliver() {
	fmt.Println("Невозможно доставить: товар еще не оплачен")
}

// Получение
func (*OrderedState) Recieve() {
	fmt.Println("Невозможно получить: товар еще не оплачен")
}

// Отмена
func (s *OrderedState) Cancel() {
	fmt.Printf("Заказ товара %s отменен\n", s.product.Title)
	s.product.Quantity++
	s.product.ChangeState(NewReadyState(s.product))
}

// Возврат
func (*OrderedState) Return() {
	fmt.Println("Невозможно вернуть: товар еще не оплачен")
}

// Структура "товар оплачен"
type PayedState struct {
	product *Product
}

// Конструктор состояния
func NewPayedState(product *Product) *PayedState {
	return &PayedState{
		product: product,
	}
}

// Заказ
func (*PayedState) Order() {
	fmt.Println("Невозможно заказать: товар уже заказан")
}

// Оплата
func (*PayedState) Pay() {
	fmt.Println("Невозможно оплатить: товар уже оплачен")
}

// Доставка
func (s *PayedState) Deliver() {
	fmt.Printf("Товар %s успешно отправлен\n", s.product.Title)
	// s.product.ChangeState(NewDeliveredState(s.product))
}

// Получение
func (*PayedState) Recieve() {
	fmt.Println("Невозможно получить: товар еще не доставлен")
}

// Отмена
func (s *PayedState) Cancel() {
	fmt.Printf("Оплата товара %s отменена, деньги возвращены: %.2f руб.\n",
		s.product.Title, s.product.Price)
	s.product.Quantity++
	s.product.ChangeState(NewReadyState(s.product))
}

// Возврат
func (*PayedState) Return() {
	fmt.Println("Невозможно вернуть: товар еще не доставлен")
}

func main() {
	fmt.Println(" \n[ СОСТОЯНИЕ ]\n ")
}
