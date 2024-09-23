package main

import (
	"fmt"
	"log"
	"os"

	ini "gopkg.in/gcfg.v1"
)

func main() {
	fmt.Println(" \n[ INI-КОНФИГУРАЦИЯ ]\n ")

	// Структура для хранения опций
	config := struct {
		Db struct {
			Login     string
			Password  int
			IsActive  bool
			Is_banned bool
		}
		Email struct {
			Admin   string
			Support string
			Noreply string
		}
	}{}

	// Чтение конфигурации
	path := os.Getenv("GOPATH") + "/src/learn/ini/"
	err := ini.ReadFileInto(&config, path+"example.ini")
	if err != nil {
		log.Fatalf("Не удалось прочитать ini-файл: %s", err)
	}

	// Секция "DB"
	fmt.Println("[DB]")
	fmt.Printf("Login: %q (%T)\n", config.Db.Login, config.Db.Login)
	fmt.Printf("Password: %d (%T)\n", config.Db.Password, config.Db.Password)
	fmt.Printf("IsActive: %v (%T)\n", config.Db.IsActive, config.Db.IsActive)
	fmt.Printf("IsBanned: %v (%T)\n", config.Db.Is_banned, config.Db.Is_banned)
	fmt.Println()

	// Секция "Email"
	fmt.Println("[Email]")
	fmt.Printf("Admin: %q (%T)\n", config.Email.Admin, config.Email.Admin)
	fmt.Printf("Support: %q (%T)\n", config.Email.Support, config.Email.Support)
	fmt.Printf("Noreply: %q (%T)\n", config.Email.Noreply, config.Email.Noreply)
}
