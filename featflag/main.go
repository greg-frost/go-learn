package main

import (
	"fmt"
)

// Флаг использования нового функционала
const useNewFeature = true

// Старый функционал
func OldFeature() {
	fmt.Println("Используется старый функционал...")
}

// Новый функционал
func NewFeature() {
	fmt.Println("Используется новый функционал...")
}

func main() {
	fmt.Println(" \n[ FEATURE FLAG ]\n ")

	// Выбор функционала
	if useNewFeature {
		NewFeature()
	} else {
		OldFeature()
	}
}
