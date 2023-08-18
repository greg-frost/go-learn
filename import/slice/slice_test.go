package slice

import (
	"reflect"
	"testing"
)

// Структура "тестовая пара суммы"
type testPairSum struct {
	values Slice
	result Element
}

// Тестовые данные для суммы
var sumTests = []testPairSum{
	{Slice{1, 2, 3}, 6},
	{Slice{5, 1, 12}, 18},
	{Slice{0, 0, 0}, 0},
	{Slice{-10, 50, -105}, -65},
	{Slice{999, 2, -1}, 1000},
}

// Стурктура "тестовая пара мэппинга"
type testPairMap struct {
	values Slice
	op     func(Element) Element
	result Slice
}

// Тестовые данные для мэппинга
var op_inc = func(v Element) Element { return v + 1 }
var op_twice = func(v Element) Element { return v * 2 }
var op_third = func(v Element) Element { return v * 3 }

var mapTests = []testPairMap{
	{Slice{1, 2, 3}, op_inc, Slice{2, 3, 4}},
	{Slice{2, 7, 0}, op_twice, Slice{4, 14, 0}},
	{Slice{1, -1, 3}, op_third, Slice{3, -3, 9}},
	{Slice{-5, -3, 10}, op_twice, Slice{-10, -6, 20}},
	{Slice{-2, -1, 0}, op_inc, Slice{-1, 0, 1}},
}

// Стурктура "тестовая пара свертки"
type testPairFold struct {
	values Slice
	op     func(Element, Element) Element
	init   Element
	result Element
}

// Тестовые данные для свертки
var op_sum = func(v1, v2 Element) Element { return v1 + v2 }
var op_prod = func(v1, v2 Element) Element { return v1 * v2 }

var foldTests = []testPairFold{
	{Slice{1, 2, 3}, op_sum, 0, 6},
	{Slice{4, 6, 2}, op_prod, 1, 48},
	{Slice{1, -1, 5}, op_sum, 10, 15},
	{Slice{2, 1, 5}, op_prod, 10, 100},
	{Slice{-10, -22, -2}, op_sum, -10, -44},
}

/* Тесты */

// Тест суммы
func TestSumSlice(t *testing.T) {
	for _, pair := range sumTests {
		g := SumSlice(pair.values)
		e := pair.result

		if g != e {
			t.Errorf("Сумма: получено %v, ожидается %v", g, e)
		}
	}
}

// Тест мэппинга
func TestMapSlice(t *testing.T) {
	for _, pair := range mapTests {
		g := make(Slice, 3)
		copy(g, pair.values)
		MapSlice(g, pair.op)
		e := pair.result

		if !reflect.DeepEqual(g, e) {
			t.Errorf("Мэппинг: получено %v, ожидается %v", g, e)
		}
	}
}

// Тест свертки
func TestFoldSlice(t *testing.T) {
	for _, pair := range foldTests {
		g := FoldSlice(pair.values, pair.op, pair.init)
		e := pair.result

		if g != e {
			t.Errorf("Свертка: получено %v, ожидается %v", g, e)
		}
	}
}
