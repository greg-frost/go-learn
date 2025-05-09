package core

// Интерфейс "регистратор транзакций"
type TransactionLogger interface {
	WritePut(key, value string)
	WriteDelete(key string)
	Err() <-chan error

	Run()
	Read() (<-chan Event, <-chan error)
}

// Структура "событие"
type Event struct {
	Sequence  uint64
	EventType EventType
	Key       string
	Value     string
}

// Тип события
type EventType byte

// Виды типов событий
const (
	_                  = iota
	EventPut EventType = iota
	EventDelete
)
