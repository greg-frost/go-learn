package main

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"go-learn/base"

	"github.com/BurntSushi/toml"
)

func main() {
	fmt.Println(" \n[ TOML-КОНФИГУРАЦИЯ ]\n ")

	// Структура для хранения опций
	config := struct {
		Title string
		User  struct {
			Id       int
			Rate     float64
			Name     string
			Birthday time.Time
		}
		Db struct {
			Server  string
			Ports   []int
			Enabled bool
		}
		Servers struct {
			Alpha struct {
				Ip string
				Dc string
			}
			Omega struct {
				Ip string
				Dc string
			}
		}
		Clients struct {
			Data  [][]string
			Hosts []string
		}
	}{}

	// Чтение конфигурации
	path := base.Dir("toml")
	if _, err := toml.DecodeFile(filepath.Join(path, "example.toml"), &config); err != nil {
		log.Fatal(err)
	}

	/* Вывод */

	fmt.Printf("Title: %v (%T)\n", config.Title, config.Title)
	fmt.Println()

	fmt.Println("User:")
	fmt.Printf("   ID: %v (%T)\n", config.User.Id, config.User.Id)
	fmt.Printf("   Rate: %v (%T)\n", config.User.Rate, config.User.Rate)
	fmt.Printf("   Name: %v (%T)\n", config.User.Name, config.User.Name)
	fmt.Printf("   Birthday: %v (%T)\n", config.User.Birthday, config.User.Birthday)
	fmt.Println()

	fmt.Println("Db:")
	fmt.Printf("   Server: %v (%T)\n", config.Db.Server, config.Db.Server)
	fmt.Printf("   Ports: %v (%T)\n", config.Db.Ports, config.Db.Ports)
	fmt.Printf("   Enabled: %v (%T)\n", config.Db.Enabled, config.Db.Enabled)
	fmt.Println()

	fmt.Println("Servers:")
	fmt.Printf("   Alpha:\n")
	fmt.Printf("      IP: %v (%T)\n", config.Servers.Alpha.Ip, config.Servers.Alpha.Ip)
	fmt.Printf("      DC: %v (%T)\n", config.Servers.Alpha.Dc, config.Servers.Alpha.Dc)
	fmt.Printf("   Omega:\n")
	fmt.Printf("      IP: %v (%T)\n", config.Servers.Omega.Ip, config.Servers.Omega.Ip)
	fmt.Printf("      DC: %v (%T)\n", config.Servers.Omega.Dc, config.Servers.Omega.Dc)
	fmt.Println()

	fmt.Println("Clients:")
	fmt.Printf("   Data: %v (%T)\n", config.Clients.Data, config.Clients.Data)
	fmt.Printf("   Hosts: %v (%T)\n", config.Clients.Hosts, config.Clients.Hosts)
}
