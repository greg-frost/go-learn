package main

import (
	"fmt"
	"math/rand"

	"github.com/spf13/viper"
)

// Включение функционала вручную
// const useNewFeature = true

// Флаг использования функционала (по переменной окружения)
// func FeatureEnabled(feature string) bool {
// 	state := strings.ToLower(os.Getenv(feature))
// 	return state == "true" || state == "on" || strings.HasPrefix(state, "enable")
// }

// Тип "включение функционала"
type Enabled func(int) bool

// Функции включения функционала
var enabledFunctions = map[string]Enabled{
	"use_new_feature": enabledByChance,
}

// Включение с определенной вероятностью
func enabledByChance(chance int) bool {
	return rand.Intn(100) < chance
}

// Флаг использования функционала (с определенной вероятностью)
func FeatureEnabled(feature string) bool {
	// Если задан флаг - использовать его
	if viper.IsSet(feature) {
		return viper.GetBool(feature)
	}

	// Поиск функции включения функционала
	enabledFunc, ok := enabledFunctions[feature]
	if !ok {
		return false
	}

	// Вызов функции включения функционала
	return enabledFunc(25)
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

	// Имя флага и регистрация в Viper
	feature := "use_new_feature"
	viper.BindEnv(feature)

	// Выбор функционала
	if FeatureEnabled(feature) {
		NewFeature()
	} else {
		OldFeature()
	}
}
