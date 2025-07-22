package main

import (
	"fmt"
)

// Интерфейс "бинарный"
type Binary interface {
	BigEndian() []byte
}

// Структура "от старшего к младшему"
type BigEndian struct {
	bytes []byte
}

// Конструктор "от старшего к младшему"
func NewBigEndian(bytes []byte) *BigEndian {
	return &BigEndian{
		bytes: bytes,
	}
}

// Байты "от старшего к младшему"
func (b *BigEndian) BigEndian() []byte {
	return b.bytes
}

// Структура "от младшего к старшему"
type LittleEndian struct {
	bytes []byte
}

// Конструктор "от младшего к старшему"
func NewLittleEndian(bytes []byte) *LittleEndian {
	return &LittleEndian{
		bytes: bytes,
	}
}

// Байты "от младшего к старшему"
func (b *LittleEndian) LittleEndian() []byte {
	return b.bytes
}

// Структура "бинарный адаптер"
type BinaryAdapter struct {
	littleEndian *LittleEndian
}

// Конструктор адаптера
func NewBinaryAdapter(littleEndian *LittleEndian) Binary {
	return &BinaryAdapter{
		littleEndian: littleEndian,
	}
}

// Байты адаптера "от старшего к младшему"
func (a *BinaryAdapter) BigEndian() []byte {
	le := a.littleEndian.LittleEndian()
	b := make([]byte, len(le))
	copy(b, le)

	left, right := 0, len(b)-1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}

	return b
}

// Структура "сеть"
type Network struct {
	binary Binary
}

// Констурктор сети
func NewNetwork(binary Binary) *Network {
	return &Network{
		binary: binary,
	}
}

// Обработка байт (big-endian)
func (n *Network) ProcessBytes() string {
	return string(n.binary.BigEndian())
}

func main() {
	fmt.Println(" \n[ АДАПТЕР ]\n ")

	// Число 21, записанное с использованием разных порядков байт
	bigEndian := NewBigEndian([]byte{'0', '0', '0', '1', '0', '1', '0', '1'})
	littleEndian := NewLittleEndian([]byte{'1', '0', '1', '0', '1', '0', '0', '0'})

	// Старшеконечный порядок
	network := NewNetwork(bigEndian)
	fmt.Println("BigEndian:", network.ProcessBytes())

	// Адаптер (младшеконечного к старшеконечному)
	adapter := NewBinaryAdapter(littleEndian)
	networkAdapted := NewNetwork(adapter)
	fmt.Println("BE Adapter:", networkAdapted.ProcessBytes())

	// Младшеконечный порядок
	fmt.Println("LittleEndian:", string(littleEndian.LittleEndian()))
}
