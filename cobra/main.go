package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// Переменные флагов
var pStr string
var pInt int
var pBool bool

// Главная команда
var rootCmd = &cobra.Command{
	Use:   "cobra",
	Short: "Главная",
	Long:  "Главная команда Cobra",
}

// Команда флагов
var flagsCmd = &cobra.Command{
	Use:   "flags",
	Short: "Флаги",
	Long:  "Команда флагов Cobra",
	Run:   flagsFunc,
}

// Обработчик флагов
func flagsFunc(cmd *cobra.Command, args []string) {
	fmt.Println("Строковый флаг:", pStr)
	fmt.Println("Числовой флаг:", pInt)
	fmt.Println("Логический флаг:", pBool)
	fmt.Println("Оставшиеся аргументы:", args)
}

func init() {
	// Настройка флагов
	flagsCmd.Flags().StringVarP(&pStr, "string", "s", "empty", "строковый флаг")
	flagsCmd.Flags().IntVarP(&pInt, "number", "n", 0, "числовой флаг")
	flagsCmd.Flags().BoolVarP(&pBool, "boolean", "b", false, "логический флаг")

	// Настройка команд
	rootCmd.AddCommand(flagsCmd)
}

func main() {
	fmt.Println(" \n[ COBRA ]\n ")

	// Запуск команд и флагов
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
