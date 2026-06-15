package optimization

import "sync"

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

// Структура "входные данные"
type Input struct {
	a int64
	b int64
}

// Структура "результат"
type Result struct {
	sumA int64
	sumB int64
}

// Структура "результат (с выравниванием)"
type ResultAligned struct {
	sumA int64
	_    [56]byte // Выравнивание
	sumB int64
}

// Подсчет сумм
func Count(inputs []Input) Result {
	var result Result
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < len(inputs); i++ {
			result.sumA += inputs[i].a
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < len(inputs); i++ {
			result.sumB += inputs[i].b
		}
	}()
	wg.Wait()
	return result
}

// Подсчет сумм (оптимизированный)
func CountOptimized(inputs []Input) ResultAligned {
	var result ResultAligned
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < len(inputs); i++ {
			result.sumA += inputs[i].a
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < len(inputs); i++ {
			result.sumB += inputs[i].b
		}
	}()
	wg.Wait()
	return result
}

// Инкремент
func Increment(s [2]int64, n int) [2]int64 {
	for i := 0; i < n; i++ {
		s[0]++
		if s[0]%2 == 0 {
			s[1]++
		}
	}
	return s
}

// Инкремент (оптимизированный)
func IncrementOptimized(s [2]int64, n int) [2]int64 {
	for i := 0; i < n; i++ {
		v := s[0]
		s[0] = v + 1
		if v%2 != 0 {
			s[1]++
		}
	}
	return s
}
