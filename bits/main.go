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

// Битовая группа
type BitPack uint32

// Размер группы
const size = 4

// Маска группы
const mask = (1 << (size + 1)) - 1

// Установка значения
func (p *BitPack) Set(i, v int) {
	p.Clear(i)
	s := (i - 1) * size
	b := BitPack(v << s)
	*p |= b
}

// Очистка значения
func (p *BitPack) Clear(i int) {
	s := (i - 1) * size
	b := BitPack(mask << s)
	*p &^= b
}

// Получение значения
func (p BitPack) Get(i int) int {
	s := (i - 1) * size
	b := int(p >> s)
	return b & mask
}

func main() {
	fmt.Println(" \n[ БИТЫ ]\n ")

	// Побитовые операции
	fmt.Println("Один бит:")
	var b Bits
	b = Set(b, Execute)    // Execute: 0 -> 1
	b = Toggle(b, Execute) // Execute: 1 -> 0
	b = Toggle(b, Execute) // Execute: 0 -> 1
	b = Set(b, Write)      // Write: 0 -> 1
	b = Clear(b, Write)    // Write: 1 -> 0
	b = Clear(b, Write)    // Write: 0 -> 0
	b = Toggle(b, Read)    // Read: 0 -> 1
	fmt.Println("Read:", Has(b, Read))
	fmt.Println("Write:", Has(b, Write))
	fmt.Println("Execute:", Has(b, Execute))
	fmt.Println()

	// Групповые побитовые операции
	fmt.Println("Группы битов:")
	var p BitPack
	fmt.Println("Установка 1 в 1")
	p.Set(1, 2)
	fmt.Println("Установка 4 в 2")
	p.Set(2, 4)
	fmt.Println("Установка 6 в 3")
	p.Set(3, 6)
	fmt.Println("Установка 2 в 1")
	p.Set(1, 2)
	fmt.Println("Получение 1:", p.Get(1))
	fmt.Println("Получение 2:", p.Get(2))
	fmt.Println("Получение 3:", p.Get(3))
	fmt.Println("Получение 4:", p.Get(4))
}
