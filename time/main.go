package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(" \n[ ВРЕМЯ ]\n ")

	// Текущее время
	now := time.Now()
	fmt.Println("Сейчас:", now)

	// Определенное время
	then := time.Date(1987, 8, 21, 12, 23, 34, 45, time.UTC)
	fmt.Println("Тогда:", then)

	// Компоненты даты и времени
	fmt.Println()
	fmt.Println("Год:", then.Year())
	fmt.Println("Месяц:", then.Month())
	fmt.Println("День:", then.Day())
	fmt.Println("День недели:", then.Weekday())
	fmt.Println("Час:", then.Hour())
	fmt.Println("Минута:", then.Minute())
	fmt.Println("Секунда:", then.Second())
	fmt.Println("Наносекунда:", then.Nanosecond())
	fmt.Println("Местоположение:", then.Location())

	// Сравнение дат
	fmt.Println()
	fmt.Println("Раньше:", then.Before(now))
	fmt.Println("Позже:", then.After(now))
	fmt.Println("Сегодня:", then.Equal(now))

	// Интервал времени
	diff := now.Sub(then)
	fmt.Println()
	fmt.Println("Интервал времени:", diff)

	// Компоненты интервала
	fmt.Println()
	fmt.Println("Лет:", diff.Hours()/(24*365))
	fmt.Println("Месяцев:", diff.Hours()/(24*30))
	fmt.Println("Дней:", diff.Hours()/24)
	fmt.Println("Часов:", diff.Hours())
	fmt.Println("Минут:", diff.Minutes())
	fmt.Println("Секунд:", diff.Seconds())
	fmt.Println("Наносекунд:", diff.Nanoseconds())

	// Прибавление/вычитание интервала
	fmt.Println()
	fmt.Println("+ интервал:", then.Add(diff))
	fmt.Println("- интервал:", then.Add(-diff))
}
