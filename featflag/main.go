package main

import (
	"fmt"
	"os"
	"strings"
)

// Включение функционала вручную
// const useNewFeature = true

// Флаг использования функционала
func FeatureEnabled(feature string) bool {
	state := strings.ToLower(os.Getenv(feature))
	return state == "true" || state == "on" || strings.HasPrefix(state, "enable")
}

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
	if FeatureEnabled("USE_NEW_FEATURE") {
		NewFeature()
	} else {
		OldFeature()
	}
}
