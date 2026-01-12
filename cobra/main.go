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
	Use:  "flags",
	Long: "Пример команды флагов для Cobra",
	Run:  flagsFunc,
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
	rootCmd.Flags().StringVarP(&pStr, "string", "s", "empty", "строковый флаг")
	rootCmd.Flags().IntVarP(&pInt, "number", "n", 0, "числовой флаг")
	rootCmd.Flags().BoolVarP(&pBool, "boolean", "b", false, "логический флаг")
}

func main() {
	fmt.Println(" \n[ COBRA ]\n ")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
