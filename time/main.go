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

	// Unix
	fmt.Println()
	fmt.Println("UNIX, с:", now.Unix())
	fmt.Println("UNIX, мс:", now.UnixMilli())
	fmt.Println("UNIX, нс:", now.UnixNano())

	/* Форматирование */

	fmt.Println()
	fmt.Println("Форматирование:")

	fmt.Println(now.Format(time.RFC3339))
	then, _ = time.Parse(time.RFC3339, "1987-08-21T12:23:34+00:00")
	fmt.Println(then)

	fmt.Println(now.Format("3:04PM"))
	fmt.Println(now.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(now.Format("2006-01-02T15:04:05.999999-07:00"))

	short, _ := time.Parse("3 04 PM", "8 41 PM")
	fmt.Println(short)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	_, err := time.Parse("Mon Jan _2 15:04:05 2006", "8:41PM")
	if err != nil {
		fmt.Println("Ошибка парсинга даты")
	}
}
