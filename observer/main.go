package main

import (
	"fmt"
)

// Интерфейс "наблюдаемый субъект"
type Observable interface {
	AddObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
	Header() string
	Body() string
	Footer() string
}

// Интерфейс "наблюдатель"
type Observer interface {
	Update()
	UpdatePush(header, body, footer string)
	UpdatePull(observable Observable)
}

// Структура "страница"
type page struct {
	observers map[Observer]bool
	header    string
	body      string
	footer    string
}

// Конструктор страницы
func NewPage() *page {
	return &page{
		observers: make(map[Observer]bool),
	}
}

// Добавление наблюдателя
func (p *page) AddObserver(observer Observer) {
	p.observers[observer] = true
}

// Удаление наблюдателя
func (p *page) RemoveObserver(observer Observer) {
	delete(p.observers, observer)
}

// Оповещение наблюдателей
func (p *page) NotifyObservers() {
	for observer := range p.observers {
		// observer.Update()
		// observer.UpdatePush(p.header, p.body, p.footer)
		observer.UpdatePull(p)
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

// Обновление состояния браузера (Push)
func (b *browser) UpdatePush(header, body, footer string) {
	fmt.Printf("Браузер: <%s> <%s> <%s>\n", header, body, footer)
}

// Обновление состояния браузера (Pull)
func (b *browser) UpdatePull(observable Observable) {
	fmt.Printf("Браузер: <%s> <%s> <%s>\n",
		observable.Header(),
		observable.Body(),
		observable.Footer(),
	)
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

// Обновление состояния логгера (Push)
func (b *logger) UpdatePush(header, body, footer string) {
	fmt.Printf("Логгер: %s, %s, %s\n", header, body, footer)
}

// Обновление состояния логгера (Pull)
func (b *logger) UpdatePull(observable Observable) {
	fmt.Printf("Логгер: %s, %s, %s\n",
		observable.Header(),
		observable.Body(),
		observable.Footer(),
	)
}

func main() {
	fmt.Println(" \n[ НАБЛЮДАТЕЛЬ ]\n ")

	// Страница
	fmt.Println("Создание и изменение страницы")
	page := NewPage()
	page.Change("Голова", "Тело", "Ноги")
	fmt.Println()

	// Браузер и логгер
	browser := NewBrowser()
	logger := NewLogger()

	// Добавление наблюдателей
	fmt.Println("Добавление браузера и логгера")
	page.AddObserver(browser)
	page.AddObserver(logger)

	fmt.Println("Изменение страницы")
	page.Change("Голова", "Тело", "Футер")
	fmt.Println()

	// Удаление логгера
	fmt.Println("Удаление логгера")
	page.RemoveObserver(logger)

	fmt.Println("Изменение страницы")
	page.Change("Header", "Body", "Footer")
	fmt.Println()

	// Удаление браузера
	fmt.Println("Удаление браузера")
	page.RemoveObserver(browser)

	fmt.Println("Изменение страницы")
	page.Change("Head", "Body", "Foot")
}
