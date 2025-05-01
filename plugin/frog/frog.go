package main

// Структура "лягушка"
type frog struct{}

// Лягушка говорит...
func (f frog) Says() string {
	return "Ква!"
}

// Символ
var Animal frog
