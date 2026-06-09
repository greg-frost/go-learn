package optimization

// Сумма всех значений
func Sum(s []int64) int64 {
	var total int64
	for i := 0; i < len(s); i++ {
		total += s[i]
	}
	return total
}

// Сумма каждого второго значения
func Sum2(s []int64) int64 {
	var total int64
	for i := 0; i < len(s); i += 2 {
		total += s[i]
	}
	return total
}

// Сумма каждого восьмого значения
func Sum8(s []int64) int64 {
	var total int64
	for i := 0; i < len(s); i += 8 {
		total += s[i]
	}
	return total
}

// Структура "пара"
type Pair struct {
	a int64
	b int64
}

// Сумма среза пар
func SumPair(pairs []Pair) int64 {
	var total int64
	for i := 0; i < len(pairs); i++ {
		total += pairs[i].a
		_ = pairs[i].b
	}
	return total
}

// Структура "пары"
type Pairs struct {
	a []int64
	b []int64
}

// Сумма пар срезов
func SumPairs(pairs Pairs) int64 {
	var total int64
	for i := 0; i < len(pairs.a); i++ {
		total += pairs.a[i]
		_ = pairs.b[i]
	}
	return total
}
