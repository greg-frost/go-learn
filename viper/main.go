package main

import (
	"fmt"
	"log"

	"go-learn/base"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	// _ "github.com/spf13/viper/remote"
)

// Путь
var path = base.Dir("viper")

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
	fmt.Println("Установка вручную")
	viper.Set("set", true)
	viper.Set("key", "value")
	fmt.Println("set:", viper.GetBool("set"))
	fmt.Printf("key: %q\n", viper.GetString("key"))
	fmt.Println()

	// Флаги командной строки
	// (приоритет ниже, чем при установке вручную)
	fmt.Println("Флаги командной строки")
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("string: %q\n", viper.GetString("string"))
	fmt.Println("number:", viper.GetInt("number"))
	fmt.Println("boolean:", viper.GetBool("boolean"))
	fmt.Println()

	// Переменные окружения
	// (приоритет ниже, чем у флагов командной строки)
	fmt.Println("Переменные окружения")
	viper.BindEnv("id")
	viper.BindEnv("port", "PORT_N")
	viper.BindEnv("string")
	fmt.Println("ID:", viper.GetInt("id"))
	fmt.Println("PORT_N:", viper.GetInt("port"))
	fmt.Printf("STRING: %q\n", viper.GetString("string"))
	fmt.Println()

	// Файлы конфигурации
	// (приоритет ниже, чем у переменных окружения)
	fmt.Println("Файлы конфигурации")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml") // Необязательно
	viper.AddConfigPath(".")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("host: %q\n", viper.GetString("host"))
	fmt.Println("port:", viper.GetInt("port"))
	fmt.Println("tags.debug:", viper.GetBool("tags.debug"))
	fmt.Println("tags.silent:", viper.GetBool("tags.silent"))
	fmt.Println("tags.ssl:", viper.GetString("tags.ssl"))
	fmt.Println()

	// Наблюдение за конфигурацией
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Конфигурация изменилась:", e.Name)
	})

	// Удаленные службы конфигурации
	// (приоритет ниже, чем у файлов конфигурации)
	// fmt.Println("Удаленные службы конфигурации")
	// viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/service.json")
	// viper.SetConfigType("json")
	// if err := viper.ReadRemoteConfig(); err != nil {
	// 	log.Fatal(err)
	// }

	// Значения по умолчанию
	// (самый низкий приоритет)
	fmt.Println("Значения по умолчанию")
	viper.SetDefault("id", 1000)
	fmt.Println("id:", viper.GetInt("id"))
}
