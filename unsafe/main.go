package main

import (
	"encoding/binary"
	"fmt"
	"math/bits"
	"unsafe"
)

// Структура "данные"
type Data struct {
	Value  uint32   // 4 байта
	Label  [10]byte // 10 байт
	Active bool     // 1 байт
	// + 1 байт для ровного счета
}

// Флаг младшеконечного порядка битов
var isLE bool

// Инициализация - проверка порядка битов
func init() {
	var x uint16 = 0xFF00
	xb := *(*[2]byte)(unsafe.Pointer(&x))
	isLE = (xb[0] == 0x00)
}

// Данные из битов (unsafe)
func DataFromBytesUnsafe(b [16]byte) Data {
	data := *(*Data)(unsafe.Pointer(&b))
	if isLE {
		data.Value = bits.ReverseBytes32(data.Value)
	}
	return data
}

// Данные из битов
func DataFromBytes(b [16]byte) Data {
	d := Data{}
	d.Value = binary.BigEndian.Uint32(b[:4])
	copy(d.Label[:], b[4:14])
	d.Active = b[14] != 0
	return d
}

// Байты из данных (unsafe)
func BytesFromDataUnsafe(d Data) [16]byte {
	if isLE {
		d.Value = bits.ReverseBytes32(d.Value)
	}
	b := *(*[16]byte)(unsafe.Pointer(&d))
	return b
}

// Байты из данных
func BytesFromData(d Data) [16]byte {
	out := [16]byte{}
	binary.BigEndian.PutUint32(out[:4], d.Value)
	copy(out[4:14], d.Label[:])
	if d.Active {
		out[14] = 1
	}
	return out
}

func main() {
	fmt.Println(" \n[ ПАКЕТ UNSAFE ]\n ")

	/* Информация */

	d := Data{
		Value:  8675309,
		Active: true,
	}
	copy(d.Label[:], "Unsafe")
	fmt.Println("Инфо:", d, unsafe.Alignof(d), unsafe.Alignof(d.Value),
		unsafe.Alignof(d.Label), unsafe.Alignof(d.Active))

	/* Из байтов в данные */

	b := [16]byte{0, 132, 95, 237, 80, 104, 111, 110, 101, 0, 0, 0, 0, 0, 1, 0}
	fmt.Println("До передачи:", b)

	b1 := BytesFromData(d)
	b2 := BytesFromDataUnsafe(d)
	if b1 != b2 {
		panic(fmt.Sprintf("%v %v", b1, b2))
	}
	fmt.Printf("Передача: %+v\n", b1)

	/* Из данных в байты */

	d1 := DataFromBytes(b1)
	d2 := DataFromBytesUnsafe(b1)
	if d1 != d2 {
		panic(fmt.Sprintf("%v %v", d1, d2))
	}
	fmt.Println("Обратно:", d1)
}
