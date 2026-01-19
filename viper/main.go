package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println(" \n[ VIPER ]\n ")

	// Установка вручную
	fmt.Println("Установка вручную:")
	viper.Set("set", true)
	viper.Set("key", "value")
	fmt.Println("set:", viper.GetBool("set"))
	fmt.Printf("key: %q\n", viper.GetString("key"))
}
