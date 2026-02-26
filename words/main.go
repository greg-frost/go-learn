package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"go-learn/base"
)

// Структура "слова"
type words struct {
	mu    sync.Mutex
	found map[string]int
}

// Конструктор слов
func NewWords() *words {
	return &words{
		found: make(map[string]int),
	}
}

// Добавление слов
func (w *words) Add(word string, n int) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.found[word] += n
}

// Подсчет слов
func CountWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		word = strings.ReplaceAll(word, ",", "")
		word = strings.ReplaceAll(word, ".", "")
		word = strings.ReplaceAll(word, "!", "")
		dict.Add(word, 1)
	}
	return scanner.Err()
}

func main() {
	fmt.Println(" \n[ ПОДСЧЕТ СЛОВ ]\n ")

	// Смена директории
	path := base.Dir("words")
	os.Chdir(path)

	if len(os.Args) == 1 {
		fmt.Println("Передайте список файлов в виде параметров!")
		return
	}

	// Параллельное сжатие
	fmt.Println("Идет подсчет слов в файлах...")
	var wg sync.WaitGroup
	w := NewWords()
	for _, file := range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			if err := CountWords(filename, w); err != nil {
				fmt.Println("Ошибка:", err)
			}
		}(file)
	}
	wg.Wait()

	// Вывод
	fmt.Println()
	for word, count := range w.found {
		if len(word) >= 3 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
}
