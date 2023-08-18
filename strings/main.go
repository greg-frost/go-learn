package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Переворот строки
func reverseString(s string) string {
	n := 0
	runes := make([]rune, len(s))
	for _, r := range s {
		runes[n] = r
		n++
	}
	runes = runes[0:n]

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

// Счетчик букв
func lettersCount(s string) map[rune]int {
	res := make(map[rune]int)

	for _, v := range s {
		res[v]++
	}

	return res
}

// Счетчик слов
func wordsCount(s string) map[string]int {
	res := make(map[string]int)
	s = strings.ToLower(s)

	words := strings.FieldsFunc(s, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})
	for _, v := range words {
		res[v]++
	}

	return res
}

// Тип "IP-адрес"
type IPAddr [4]byte

// Стрингер для IP-адресов
func (octs IPAddr) String() string {
	res := make([]string, len(octs))
	for i, v := range octs {
		res[i] = fmt.Sprintf("%3v", v)
	}
	return strings.Join(res, ".")
}

// Обрезка строки
func trim(text string) string {
	if text == "" {
		return text
	}

	if text[0] != ' ' && text[len(text)-1] != ' ' {
		return text
	}

	start := 0
	end := len(text)

	for start < end {
		if text[start] != ' ' {
			break
		}
		start++
	}

	for end > start {
		if text[end-1] != ' ' {
			break
		}
		end--
	}

	return text[start:end]
}

func main() {
	fmt.Println(" \n[ СТРОКИ ]\n ")

	var str string = "Hello" + ", " + "Мир!"

	fmt.Println(`"` + str + `"`)
	fmt.Println()

	/* Длина и символ */

	fmt.Println("Длина:", len(str))
	fmt.Println("Первый символ:", str[0])

	fmt.Println()

	/* Подстрока, поиск и позиция */

	fmt.Println("Подстрока \"Мир\":", strings.Contains(str, "Мир"))
	fmt.Println("Число букв \"L\":", strings.Count(str, "l"))
	fmt.Println("Позиция буквы \"L\":", strings.Index(str, "l"))

	fmt.Println()

	/* Замена, префикс и суффикс */

	fmt.Println("Замена слова:", strings.Replace(str, "Мир", "World", 1))
	fmt.Println("Начинается на \"Hel\":", strings.HasPrefix(str, "Hel"))
	fmt.Println("Заканчивается на \"!?\":", strings.HasSuffix(str, "!?"))

	fmt.Println()

	/* Разбиение, слияние, повторение */

	split := strings.Split(str, ", ")
	fmt.Println("Разбиение по запятой:", split)
	fmt.Println("Слияние по запятой:", strings.Join(split, ", "))
	fmt.Println("Повторение дважды:", strings.Repeat(split[0], 2))

	fmt.Println()

	/* Регистры */

	fmt.Println("Нижний регистр:", strings.ToLower(str))
	fmt.Println("Верхний регистр:", strings.ToUpper(str))

	fmt.Println()

	/* Байты, руны и строки */

	bytes := []byte(str)
	firstBytes := make([]byte, 5)
	copy(firstBytes, bytes)
	strAgain := string(firstBytes)

	fmt.Println("Перевод в байты:", firstBytes)
	fmt.Println("Снова строка:", strAgain)

	fmt.Println()

	jap := "椒, hello, 椒!"

	fmt.Println("Длина", jap, "в строке:", len(jap))
	fmt.Println("Длина", jap, "в байтах:", len([]byte(jap)))
	fmt.Println("Длина", jap, "в рунах:", len([]rune(jap)))

	japCount := 0
	for range jap {
		japCount++
	}
	fmt.Println("Длина", jap, "в range:", japCount)

	fmt.Println("Перевернутая строка:", reverseString(jap))

	fmt.Println()

	/* Счетчики букв и слов */

	str = "АБВГДБА!"
	fmt.Print("Счетчик букв для \"", str, "\":\n")
	fmt.Println(lettersCount(str))
	fmt.Println()

	str = "Мир, Труд, Май, мир!"
	fmt.Print("Счетчик слов для \"", str, "\":\n")
	fmt.Println(wordsCount(str))
	fmt.Println()

	/* IP-адреса */

	fmt.Println("IP-адреса:")
	hosts := map[string]IPAddr{
		"Loopback":  {127, 0, 0, 1},
		"GoogleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%10v: %v\n", name, ip)
	}
}
