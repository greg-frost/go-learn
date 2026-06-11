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

// Структура "узел"
type Node struct {
	Value int64
	Next  *Node
}

// Сумма значений связного списка
func SumLinkedList(curr *Node) int64 {
	var total int64
	for curr != nil {
		total += curr.Value
		curr = curr.Next
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

// Сумма первых значений рядов (512 столбцов)
func SumRows512(s [][512]int64) int64 {
	var total int64
	for i := 0; i < len(s); i++ {
		for j := 0; j < 8; j++ {
			total += s[i][j]
		}
	}
	return total
}

// Сумма первых значений рядов (513 столбцов)
func SumRows513(s [][513]int64) int64 {
	var total int64
	for i := 0; i < len(s); i++ {
		for j := 0; j < 8; j++ {
			total += s[i][j]
		}
	}
	return total
}
