package main

import (
	"testing"
	"unicode/utf8"
)

// Обычный тест
func TestReverse(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}

	for _, tt := range tests {
		rev, _ := Reverse(tt.in)
		if rev != tt.want {
			t.Errorf("Reverse: %q, ожидается %q", rev, tt.want)
		}
	}
}

// Фаззинг-тест (генерация случайных данных)
func FuzzReverse(f *testing.F) {
	tests := []string{"Hello, world", " ", "!12345"}
	for _, tt := range tests {
		f.Add(tt)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		revrev, err2 := Reverse(rev)
		if err2 != nil {
			t.Skip()
		}

		if orig != revrev {
			t.Errorf("До: %q, После: %q", orig, revrev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse произвел неверную UTF-8 строку %q", rev)
		}
	})
}
