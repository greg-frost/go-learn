package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	fmt.Println(" \n[ YAML-КОНФИГУРАЦИЯ 2 ]\n ")

	// Структура для хранения опций
	config := struct {
		Int       int
		Oct       int
		Hex       int
		Float     float64
		Exp       float64
		Bool      bool
		Inf       float64
		Neginf    float64
		Nan       float64
		Null      interface{}
		Str       string
		Strings   string
		Paragraph string
		List      []string
		Object    struct {
			User struct {
				Login    string
				Password string
				Langs    []string
			}
			Status struct {
				Admin  bool
				Active bool
				Banned bool
			}
		}
	}{}

	// Чтение конфигурации
	path := os.Getenv("GOPATH") + "/src/golearn/yaml2/"
	source, err := ioutil.ReadFile(path + "example.yml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		log.Fatal(err)
	}

	// Секция "Числа"
	fmt.Println("[ Числа ]")
	fmt.Printf("Int: %v (%T)\n", config.Int, config.Int)
	fmt.Printf("Oct: %v (%T)\n", config.Oct, config.Oct)
	fmt.Printf("Hex: %v (%T)\n", config.Hex, config.Hex)
	fmt.Printf("Float: %v (%T)\n", config.Float, config.Float)
	fmt.Printf("Exp: %v (%T)\n", config.Exp, config.Exp)
	fmt.Printf("Bool: %v (%T)\n", config.Bool, config.Bool)
	fmt.Printf("+Inf: %v (%T)\n", config.Inf, config.Inf)
	fmt.Printf("-Inf: %v (%T)\n", config.Neginf, config.Neginf)
	fmt.Printf("NAN: %v (%T)\n", config.Nan, config.Nan)
	fmt.Printf("Null: %v (%T)\n", config.Null, config.Null)
	fmt.Println()

	// // Секция "Строки"
	fmt.Println("[ Строки ]")
	fmt.Printf("Str: %q (%T)\n", config.Str, config.Str)
	fmt.Printf("Strings: %q (%T)\n", config.Strings, config.Strings)
	fmt.Printf("Paragraph: %q (%T)\n", config.Paragraph, config.Paragraph)
	fmt.Println()

	// // Секция "Списки"
	fmt.Println("[ Списки ]")
	fmt.Printf("Элементы: %v (%T)\n", config.List, config.List)
	fmt.Printf("Количество: %d\n", len(config.List))
	fmt.Printf("Первый: %v\n", config.List[0])
	fmt.Println()

	// // Секция "Объекты"
	fmt.Println("[ Объекты ]")
	objectUser := config.Object.User
	objectStatus := config.Object.Status
	fmt.Println("User:")
	fmt.Printf("   Login: %v (%T)\n", objectUser.Login, objectUser.Login)
	fmt.Printf("   Password: %v (%T)\n", objectUser.Password, objectUser.Password)
	fmt.Printf("   Langs: %v (%T)\n", objectUser.Langs, objectUser.Langs)
	fmt.Println("Status:")
	fmt.Printf("   Admin: %v (%T)\n", objectStatus.Admin, objectStatus.Admin)
	fmt.Printf("   Active: %v (%T)\n", objectStatus.Active, objectStatus.Active)
	fmt.Printf("   Banned: %v (%T)\n", objectStatus.Banned, objectStatus.Banned)
}
