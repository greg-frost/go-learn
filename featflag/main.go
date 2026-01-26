package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// Включение функционала вручную
// const useNewFeature = true

// Флаг использования функционала
func FeatureEnabled(feature string) bool {
	// state := strings.ToLower(os.Getenv(feature))
	// return state == "true" || state == "on" || strings.HasPrefix(state, "enable")
	return viper.GetBool(feature)
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
