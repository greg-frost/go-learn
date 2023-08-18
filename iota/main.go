package main

import (
	"fmt"
)

// Тип "день недели"
type Weekday int

// Дни недели
const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// Следующий день
func NextDay(day Weekday) Weekday {
	return (day % 7) + 1
}

// Нечетные числа
const (
	one = iota*2 + 1
	three
	five
	seven
	nine
)

// Тип "размер в байтах"
type ByteSize float64

// Единицы измерения
const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// Вывод размера в нужных единицах
func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

func main() {
	fmt.Println(" \n[ ЙОТА ]\n ")

	var today Weekday = Sunday
	tomorrow := NextDay(today)
	fmt.Println("Сегодня:", today, "Завтра:", tomorrow)

	fmt.Println("Нечетные:", one, three, five, seven, nine)

	var size ByteSize = 1000000000
	fmt.Println("Размер файла:", size)
}
