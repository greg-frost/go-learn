package main

// Структура "утка"
type duck struct{}

// Утка говорит...
func (d duck) Says() string {
	return "Кря!"
}

// Символ
var Animal duck
