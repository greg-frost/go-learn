package main

import (
	"fmt"
	"log"
)

// Интерфейс "оплата"
type Payment interface {
	Pay() error
}

// Структура "оплата картой"
type cardPayment struct {
	number int
	cvv    int
}

// Конструктор оплаты картой
func NewCardPayment(number, cvv int) Payment {
	return &cardPayment{
		number: number,
		cvv:    cvv,
	}
}

// Оплата картой
func (p *cardPayment) Pay() error {
	fmt.Println("Оплата картой...")
	return nil
}

// Структура "оплата наличными"
type cashPayment struct {
	money map[int]int
}

// Конструктор оплаты наличными
func NewCashPayment(money map[int]int) Payment {
	return &cashPayment{
		money: money,
	}
}

// Оплата наличными
func (p *cashPayment) Pay() error {
	fmt.Println("Оплата наличными...")
	return nil
}

// Обработка заказа
func processOrder(product string, payment Payment) error {
	fmt.Println("Обработка заказа...")

	err := payment.Pay()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println(" \n[ СТРАТЕГИЯ ]\n ")

	product := "Товар"
	payWay := "cash"

	// Тип оплаты
	var payment Payment
	switch payWay {
	case "card":
		payment = NewCardPayment(1234567890, 123)
	case "cash":
		payment = NewCashPayment(map[int]int{5000: 1, 1000: 2, 500: 1})
	default:
		log.Fatalf("нет такого типа оплаты: %s", payWay)
	}

	// Обработка заказа
	err := processOrder(product, payment)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("OK!")
}
