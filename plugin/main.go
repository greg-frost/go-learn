package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"

	"go-learn/base"
)

// Интерфейс "говорун"
type Sayer interface {
	Says() string
}

func main() {
	fmt.Println(" \n[ ПЛАГИНЫ ]\n ")

	// Путь к плагину
	if len(os.Args) != 2 {
		fmt.Printf("Синтаксис: ./%s ANIMAL\n", filepath.Base(os.Args[0]))
		return
	}
	name := os.Args[1]
	module := fmt.Sprintf("%s/%s/%s.so", base.Dir("plugin"), name, name)

	// Открытие плагина
	p, err := plugin.Open(module)
	if err != nil {
		log.Fatal(err)
	}

	// Поиск экспортированного символа
	symbol, err := p.Lookup("Animal")
	if err != nil {
		log.Fatal(err)
	}

	// Проверка соответствия интерфейсу
	animal, ok := symbol.(Sayer)
	if !ok {
		log.Fatal("Это не говорун!")
	}

	fmt.Printf("%s говорит: %q\n", name, animal.Says())
}
