package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

// Структура "растение"
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	ID      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

// Стрингер
func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.ID, p.Name, p.Origin)
}

// Структура "вложение"
type Nesting struct {
	XMLName xml.Name `xml:"nesting"`
	Plants  []*Plant `xml:"parent>child>plant"`
}

func main() {
	fmt.Println(" \n[ XML ]\n ")

	// Структура
	coffee := &Plant{
		ID:     27,
		Name:   "Coffee",
		Origin: []string{"Ethiopia", "Brazil"},
	}

	// Маршаллинг
	out, _ := xml.MarshalIndent(coffee, "", "    ")
	fmt.Println("Маршаллинг в XML:")
	fmt.Println(xml.Header + string(out))
	fmt.Println()

	// Демаршаллинг
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Демаршаллинг из XML:")
	fmt.Println(p)
	fmt.Println()

	// Вложение
	tomato := &Plant{
		ID:     81,
		Name:   "Tomato",
		Origin: []string{"Mexico", "California"},
	}
	nesting := &Nesting{
		Plants: []*Plant{coffee, tomato},
	}
	fmt.Println("Вложение:")
	out, _ = xml.MarshalIndent(nesting, "", "    ")
	fmt.Println(string(out))
}
