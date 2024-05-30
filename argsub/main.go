package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(" \n[ АРГУМЕНТЫ С ПОДКОМАНДАМИ ]\n ")

	// Логгер
	loggerCmd := flag.NewFlagSet("logger", flag.ExitOnError)
	loggerEnable := loggerCmd.Bool("enable", false, "включение логгера")
	loggerName := loggerCmd.String("prefix", "", "префикс логгера")

	// Хранилище
	storageCmd := flag.NewFlagSet("storage", flag.ExitOnError)
	storagePlace := storageCmd.String("place", "", "место хранения")
	storageLimit := storageCmd.Int("limit", 0, "лимит памяти")

	// Если команда не задана
	if len(os.Args) < 2 {
		log.Fatal("Ожидаются команды 'logger' или 'storage'")
	}

	// Выбор команды
	switch os.Args[1] {
	// Логгер
	case "logger":
		loggerCmd.Parse(os.Args[2:])
		fmt.Println("Команда 'logger':")
		fmt.Println("   enable:", *loggerEnable)
		fmt.Println("   prefix:", *loggerName)
		fmt.Println("   остальные:", loggerCmd.Args())
	// Хранилище
	case "storage":
		storageCmd.Parse(os.Args[2:])
		fmt.Println("Команда 'storage':")
		fmt.Println("   place:", *storagePlace)
		fmt.Println("   limit:", *storageLimit)
		fmt.Println("   остальные:", storageCmd.Args())
	// Неизвестная команда
	default:
		log.Fatal("Ошибка: ожидаются команды 'logger' или 'storage'")
	}
}
