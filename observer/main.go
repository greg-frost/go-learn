package main

import (
	"fmt"
)

// Интерфейс "наблюдаемый субъект"
type Observable interface {
	AddObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

// Интерфейс "наблюдатель"
type Observer interface {
	Update()
}

// Структура "страница"
type page struct {
	observers []Observer
	header    string
	body      string
	footer    string
}

// Конструктор страницы
func NewPage() *page {
	return new(page)
}

// Добавление наблюдателя
func (p *page) AddObserver(observer Observer) {
	p.observers = append(p.observers, observer)
}

// Удаление наблюдателя
func (p *page) RemoveObserver(observer Observer) {
	observers := make([]Observer, 0, len(p.observers))
	for _, o := range p.observers {
		if o != observer {
			observers = append(observers, o)
		}
	}
	p.observers = observers
}

// Оповещение наблюдателей
func (p *page) NotifyObservers() {
	for _, observer := range p.observers {
		observer.Update()
	}
}

// Изменение страницы
func (p *page) Change(header, body, footer string) {
	p.header = header
	p.body = body
	p.footer = footer

	p.NotifyObservers() // Оповещение наблюдятелей
}

// Получение заголовка страницы
func (p *page) Header() string {
	return p.header
}

// Получение тела страницы
func (p *page) Body() string {
	return p.body
}

// Получение футера страницы
func (p *page) Footer() string {
	return p.footer
}

// Структура "браузер"
type browser struct{}

// Конструктор браузера
func NewBrowser() *browser {
	return new(browser)
}

// Обновление состояния браузера
func (b *browser) Update() {
	fmt.Println("Браузер: Рендеринг страницы")
}

// Структура "логгер"
type logger struct{}

// Конструктор логгера
func NewLogger() *logger {
	return new(logger)
}

// Обновление состояния логгера
func (l *logger) Update() {
	fmt.Println("Логгер: Логирование страницы")
}

func main() {
	fmt.Println(" \n[ НАБЛЮДАТЕЛЬ ]\n ")

	// Страница, бразузер и логгер
	page := NewPage()
	browser := NewBrowser()
	logger := NewLogger()

	// Добавление наблюдателей
	fmt.Println("Добавление браузера и логгера...")
	page.AddObserver(browser)
	page.AddObserver(logger)

	// Изменение страницы
	page.Change("Голова", "Тело", "Ноги")
	fmt.Println()

	// Удаление наблюдателя
	fmt.Println("Удаление логгера...")
	page.RemoveObserver(logger)

	// Очередное изменение страницы
	page.Change("Header", "Body", "Footer")
}
