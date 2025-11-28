package filter

import "reflect"

// Фильтрация любых значений (рефлексия)
func Filter(slice interface{}, filter interface{}) interface{} {
	sv := reflect.ValueOf(slice)
	fv := reflect.ValueOf(filter)

	sliceLen := sv.Len()
	out := reflect.MakeSlice(sv.Type(), 0, sliceLen)
	for i := 0; i < sliceLen; i++ {
		curVal := sv.Index(i)
		values := fv.Call([]reflect.Value{curVal})
		if values[0].Bool() {
			out = reflect.Append(out, curVal)
		}
	}

	return out.Interface()
}

// Фильтрация строк (без рефлексии)
func FilterString(s []string, f func(string) bool) []string {
	out := make([]string, 0, len(s))
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

// Фильтрация чисел (без рефлексии)
func FilterInt(s []int, f func(int) bool) []int {
	out := make([]int, 0, len(s))
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}
