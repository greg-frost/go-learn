package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Команда Cobra
var rootCmd = &cobra.Command{
	Use: "cobra",
}

func init() {
	// Настройка флагов в Cobra
	rootCmd.Flags().StringP("string", "s", "empty", "строковый флаг")
	rootCmd.Flags().IntP("number", "n", 0, "числовой флаг")
	rootCmd.Flags().BoolP("boolean", "b", false, "логический флаг")

	// Привязка к Viper
	viper.BindPFlag("string", rootCmd.Flags().Lookup("string"))
	viper.BindPFlag("number", rootCmd.Flags().Lookup("number"))
	viper.BindPFlag("boolean", rootCmd.Flags().Lookup("boolean"))
}

func main() {
	fmt.Println(" \n[ VIPER ]\n ")

	// Установка вручную
	// (наивысший приоритет)
	fmt.Println("Установка вручную:")
	viper.Set("set", true)
	viper.Set("key", "value")
	fmt.Println("set:", viper.GetBool("set"))
	fmt.Printf("key: %q\n", viper.GetString("key"))
	fmt.Println()

	// Флаги командной строки
	// (приоритет ниже, чем при установке вручную)
	fmt.Println("Флаги командной строки:")
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("string: %q\n", viper.GetString("string"))
	fmt.Println("number:", viper.GetInt("number"))
	fmt.Println("boolean:", viper.GetBool("boolean"))
}
