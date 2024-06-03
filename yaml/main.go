package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kylelemons/go-gypsy/yaml"
)

// Преобразование узла в карту
func nodeToMap(node yaml.Node) yaml.Map {
	m, ok := node.(yaml.Map)
	if !ok {
		log.Fatalf("%v не карта", node)
	}
	return m
}

// Преобразование узла в список
func nodeToList(node yaml.Node) yaml.List {
	m, ok := node.(yaml.List)
	if !ok {
		log.Fatalf("%v не список", node)
	}
	return m
}

func main() {
	fmt.Println(" \n[ YAML-КОНФИГУРАЦИЯ ]\n ")

	// Чтение конфигурации
	path := os.Getenv("GOPATH") + "/src/golearn/yaml/"
	config, err := yaml.ReadFile(path + "example.yaml")
	if err != nil {
		log.Fatalf("Не удалось прочитать yaml-файл: %s", err)
	}

	// Секция "Числа"
	fmt.Println("[ Числа ]")
	integer, _ := config.GetInt("int")
	fmt.Printf("Int: %d (%T)\n", integer, integer)
	oct, _ := config.Get("oct")
	fmt.Printf("Oct: %s (%T)\n", oct, oct)
	hex, _ := config.Get("hex")
	fmt.Printf("Hex: %s (%T)\n", hex, hex)
	float, _ := config.Get("float")
	fmt.Printf("Float: %s (%T)\n", float, float)
	exp, _ := config.Get("exp")
	fmt.Printf("Exp: %s (%T)\n", exp, exp)
	boolean, _ := config.GetBool("bool")
	fmt.Printf("Bool: %t (%T)\n", boolean, boolean)
	inf, _ := config.Get("inf")
	fmt.Printf("+Inf: %s (%T)\n", inf, inf)
	neginf, _ := config.Get("neginf")
	fmt.Printf("-Inf: %q (%T)\n", neginf, neginf)
	nan, _ := config.Get("nan")
	fmt.Printf("NaN: %s (%T)\n", nan, nan)
	null, _ := config.Get("null")
	fmt.Printf("Null: %s (%T)\n", null, null)
	fmt.Println()

	// Секция "Строки"
	fmt.Println("[ Строки ]")
	str, _ := config.Get("str")
	fmt.Printf("Str: %q (%T)\n", str, str)
	strings, _ := config.Get("strings")
	fmt.Printf("Strings: %q (%T)\n", strings, strings)
	paragraph, _ := config.Get("paragraph")
	fmt.Printf("Paragraph: %q (%T)\n", paragraph, paragraph)
	fmt.Println()

	// Секция "Списки"
	fmt.Println("[ Списки ]")
	list := nodeToMap(config.Root)["list"]
	listCount, _ := config.Count("list")
	fmt.Printf("Элементы: %v (%T)\n", list, list)
	fmt.Printf("Количество: %d\n", listCount)
	fmt.Printf("Первый: %v\n", nodeToList(list)[0])
	fmt.Println()

	// Секция "Объекты"
	fmt.Println("[ Объекты ]")
	object := nodeToMap(nodeToMap(config.Root)["object"])
	objectUser := nodeToMap(object["user"])
	objectStatus := nodeToMap(object["status"])
	fmt.Println("User:")
	fmt.Printf("   Login: %v (%T)\n", objectUser["login"], objectUser["login"])
	fmt.Printf("   Password: %v (%T)\n", objectUser["password"], objectUser["password"])
	fmt.Printf("   Langs: %v (%T)\n", objectUser["langs"], objectUser["langs"])
	fmt.Println("Status:")
	fmt.Printf("   Admin: %v (%T)\n", objectStatus["admin"], objectStatus["admin"])
	fmt.Printf("   Active: %v (%T)\n", objectStatus["active"], objectStatus["active"])
	fmt.Printf("   Banned: %v (%T)\n", objectStatus["banned"], objectStatus["banned"])
}
