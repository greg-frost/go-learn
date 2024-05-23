package main

import (
	"fmt"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct {
	URL    string
	Status int
}

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

	const workers = 3

	jobs := make(chan Site, workers)
	results := make(chan Result, workers)

	for w := 1; w <= workers; w++ {
		go crawl(w, jobs, results)
	}

	urls := []string{
		"https://yandex.com",
		"https://google.com",
		"https://go.dev",
		"https://go.dev/doc",
		"https://go.dev/blogs",
	}

	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)

	for r := 1; r <= len(urls); r++ {
		result := <-results
		fmt.Println("Ответ:", result)
	}
}
