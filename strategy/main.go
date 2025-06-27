package main

import (
	"fmt"
	"log"
	"strconv"
)

// Интерфейс "оплата"
type Payment interface {
	Pay(amount int) error
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
func (p *cardPayment) Pay(amount int) error {
	card := []byte(strconv.Itoa(p.number))
	for i := 0; i < len(card)-4; i++ {
		card[i] = '*'
	}
	fmt.Printf("Оплата картой: %s\n", string(card))
	return nil
}

// Структура "оплата наличными"
type cashPayment struct {
	passport string
}

// Конструктор оплаты наличными
func NewCashPayment(passport string) Payment {
	return &cashPayment{
		passport: passport,
	}
}

// Оплата наличными
func (p *cashPayment) Pay(amount int) error {
	fmt.Printf("Оплата наличными: %d руб.\n", amount)
	return nil
}

// Структура "обработчик заказов"
type orderProcessor struct {
	payment Payment
}

// Конструктор обработчика заказов
func NewOrderProcessor(payment Payment) *orderProcessor {
	return &orderProcessor{
		payment: payment,
	}
}

// Обработка заказа
func (op *orderProcessor) processOrder(product string, amount int) error {
	fmt.Println("Обработка заказа:", product)
	err := op.payment.Pay(amount)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println(" \n[ СТРАТЕГИЯ ]\n ")

	const (
		product = "Товар"
		payWay  = "card"
		amount  = 1000
	)

	// Выбор типа оплаты
	var payment Payment
	switch payWay {
	case "card":
		payment = NewCardPayment(1234567890, 123)
	case "cash":
		payment = NewCashPayment("1234-567890")
	default:
		log.Fatalf("нет такого типа оплаты: %s", payWay)
	}

	// Обработка заказа
	op := NewOrderProcessor(payment)
	err := op.processOrder(product, amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("OK: Оплата прошла успешно!")
}
