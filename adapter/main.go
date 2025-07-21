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
	b := a.littleEndian.LittleEndian()
	left, right := 0, len(b)
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
	return b
}

func main() {
	fmt.Println(" \n[ АДАПТЕР ]\n ")
}
