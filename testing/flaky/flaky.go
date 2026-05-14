package flaky

import "fmt"

// Структура "значение"
type Value struct {
	name  string
	value int
}

// Получение значений
func getValues(n int) []Value {
	res := make([]Value, 0, n)
	for i := 1; i <= n; i++ {
		res = append(res, Value{
			name:  fmt.Sprintf("Value_%d", i),
			value: i * 1000},
		)
	}
	return res
}

// Интерфейс "публикатор"
type Publisher interface {
	Publish([]Value)
}

// Структура "обработчик"
type Handler struct {
	n         int
	publisher Publisher
}

// Получение лучшего значения
func (h Handler) GetBestValue(input int) Value {
	values := getValues(input)
	best := values[0]
	go func() {
		if len(values) > h.n {
			values = values[:h.n]
		}
		h.publisher.Publish(values)
	}()
	return best
}
