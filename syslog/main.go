package main

import (
	"fmt"
)

func main() {
	fmt.Println(" \n[ СИСТЕМНОЕ ЛОГИРОВАНИЕ ]\n ")

	// Стандартный
	fmt.Println("Стандартный логгер:")
	fmt.Println("(не работает в Windows)")

	// priority := syslog.LOG_LOCAL3 | syslog.LOG_NOTICE
	// flags := log.LstdFlags | log.Lshortfile
	// logger, err := syslog.NewLogger(priority, flags)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// logger.Printf("Hello\n")
	// logger.Println("World")

	fmt.Println()

	// Уровневый
	fmt.Println("Уровневый логгер:")
	fmt.Println("(не работает в Windows)")

	// logger, err = syslog.New(syslog.LOG_LOCAL3, "syslog")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer logger.Close()
	// logger.Debug("Сообщение DEBUG")
	// logger.Notice("Сообщение NOTICE")
	// logger.Warning("Сообщение WARNING")
	// logger.Alert("Сообщение ALERT")
}
