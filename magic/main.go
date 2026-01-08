package main

import (
	"fmt"
	"reflect"
)

// Интерфейс "заклинание"
type Spell interface {
	Name() string
	Char() string
	Value() int
}

// Интерфейс "приемник заклинания"
type CastReceiver interface {
	ReceiveSpell(s Spell)
}

// Каст заклинания на всех
func CastToAll(spell Spell, objects []interface{}) {
	for _, obj := range objects {
		CastTo(spell, obj)
		fmt.Printf("Каст %s на %#v\n", spell.Name(), obj)
	}
}

// Каст заклинания на кого-то конкретного
func CastTo(spell Spell, object interface{}) {
	// Если есть метод
	if recv, ok := object.(CastReceiver); ok {
		recv.ReceiveSpell(spell)
		return
	}

	val := reflect.ValueOf(object)

	// Если не указатель на структуру
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return
	}

	field := val.Elem().FieldByName(spell.Char())

	// Если не найдено
	if !field.IsValid() {
		return
	}
	// Если нельзя изменить
	if !field.CanSet() {
		return
	}
	// Если не целое
	if field.Kind() != reflect.Int && field.Kind() != reflect.Int8 &&
		field.Kind() != reflect.Int16 && field.Kind() != reflect.Int32 &&
		field.Kind() != reflect.Int64 {
		return
	}

	field.SetInt(field.Int() + int64(spell.Value())) // Новое значение
}

// Структура "заклинание"
type spell struct {
	name string
	char string
	val  int
}

// Новое заклинание
func newSpell(name string, char string, val int) Spell {
	return &spell{name: name, char: char, val: val}
}

// Название заклинания
func (s spell) Name() string {
	return s.name
}

// Характеристика, на которую влияет заклинание
func (s spell) Char() string {
	return s.char
}

// Значение влияния заклинания
func (s spell) Value() int {
	return s.val
}

// Структура "игрок"
type Player struct {
	name   string
	health int
}

// Получение заклинания
func (p *Player) ReceiveSpell(s Spell) {
	if s.Char() == "Health" {
		p.health += s.Value() / 10
	}
}

// Структура "зомби"
type Zombie struct {
	Health int
}

// Структура "демон"
type Daemon struct {
	Health int
}

// Структура "орк"
type Orc struct {
	Health int
}

// Структура "стена"
type Wall struct {
	Durability int
}

func main() {
	fmt.Println(" \n[ МАГИЯ РЕФЛЕКСИИ ]\n ")

	// Игрок
	player := &Player{
		name:   "Player",
		health: 100,
	}

	// Враги
	enemies := []interface{}{
		&Zombie{Health: 1000},
		&Zombie{Health: 1000},
		&Orc{Health: 500},
		&Orc{Health: 500},
		&Orc{Health: 500},
		&Daemon{Health: 750},
		&Daemon{Health: 750},
		&Wall{Durability: 100},
	}

	// Заклинания
	CastToAll(newSpell("fire", "Health", -250), append(enemies, player))
	fmt.Println()
	CastToAll(newSpell("heal", "Health", 100), append(enemies, player))
}
