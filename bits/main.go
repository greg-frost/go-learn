package main

import (
	"fmt"
)

// Тип "биты"
type Bits uint8

// Константы операций
const (
	Execute Bits = 1 << iota
	Write
	Read
)

// Установка флага
func Set(b, flag Bits) Bits {
	return b | flag
}

// Выключение флага
func Clear(b, flag Bits) Bits {
	return b &^ flag
}

// Переключение флага
func Toggle(b, flag Bits) Bits {
	return b ^ flag
}

// Проверка флага
func Has(b, flag Bits) bool {
	return b&flag != 0
}

func main() {
	fmt.Println(" \n[ БИТЫ ]\n ")

	var b Bits

	// Действия
	b = Set(b, Execute)    // Execute: 0 -> 1
	b = Toggle(b, Execute) // Execute: 1 -> 0
	b = Toggle(b, Execute) // Execute: 0 -> 1
	b = Set(b, Write)      // Write: 0 -> 1
	b = Clear(b, Write)    // Write: 1 -> 0
	b = Clear(b, Write)    // Write: 0 -> 0
	b = Toggle(b, Read)    // Read: 0 -> 1

	// Вывод
	fmt.Println("Read:", Has(b, Read))
	fmt.Println("Write:", Has(b, Write))
	fmt.Println("Execute:", Has(b, Execute))
}
