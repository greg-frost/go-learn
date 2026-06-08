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
