package main

import (
	"fmt"
	"net/http"
)

// Структура "Сайт"
type Site struct {
	URL string
}

// Структура "Результат"
type Result struct {
	URL    string
	Status int
}

// Получение HTTP статуса ответа
func crawl(worker int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs {
		fmt.Printf("Воркер #%d: %s\n", worker, site.URL)

		resp, err := http.Get(site.URL)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		results <- Result{URL: site.URL, Status: resp.StatusCode}
	}
}

func main() {
	fmt.Println(" \n[ ВОРКЕРЫ ]\n ")

	// Число воркеров
	const workers = 3

	// Каналы заданий и результатов
	jobs := make(chan Site, workers)
	results := make(chan Result, workers)

	// Запуск воркеров
	for w := 1; w <= workers; w++ {
		go crawl(w, jobs, results)
	}

	// Список заданий
	urls := []string{
		"https://yandex.com",
		"https://google.com",
		"https://go.dev",
		"https://go.dev/doc",
		"https://go.dev/blogs",
	}

	// Раздача заданий
	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)

	// Получение результатов
	for r := 1; r <= len(urls); r++ {
		result := <-results
		fmt.Println("Ответ:", result)
	}
}
