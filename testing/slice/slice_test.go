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
var opInc = func(v Element) Element { return v + 1 }
var opTwice = func(v Element) Element { return v * 2 }
var opThird = func(v Element) Element { return v * 3 }

var mapTests = []testPairMap{
	{Slice{1, 2, 3}, opInc, Slice{2, 3, 4}},
	{Slice{2, 7, 0}, opTwice, Slice{4, 14, 0}},
	{Slice{1, -1, 3}, opThird, Slice{3, -3, 9}},
	{Slice{-5, -3, 10}, opTwice, Slice{-10, -6, 20}},
	{Slice{-2, -1, 0}, opInc, Slice{-1, 0, 1}},
}

// Стурктура "тестовая пара свертки"
type testPairFold struct {
	values Slice
	op     func(Element, Element) Element
	init   Element
	result Element
}

// Тестовые данные для свертки
var opSum = func(v1, v2 Element) Element { return v1 + v2 }
var opProd = func(v1, v2 Element) Element { return v1 * v2 }

var foldTests = []testPairFold{
	{Slice{1, 2, 3}, opSum, 0, 6},
	{Slice{4, 6, 2}, opProd, 1, 48},
	{Slice{1, -1, 5}, opSum, 10, 15},
	{Slice{2, 1, 5}, opProd, 10, 100},
	{Slice{-10, -22, -2}, opSum, -10, -44},
}

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
