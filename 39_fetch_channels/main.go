package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// Запросы к сайтам через каналы + структуры для обработки ответов

type Response struct {
	URL        string
	StatusCode int
	Body       string
	Error      error
	Duration   time.Duration
}

func fetchURL(url string, wg *sync.WaitGroup, results chan<- Response) {
	defer wg.Done()

	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		results <- Response{
			URL:      url,
			Error:    err,
			Duration: time.Since(start),
		}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		results <- Response{
			URL:        url,
			StatusCode: resp.StatusCode,
			Error:      err,
			Duration:   time.Since(start),
		}
		return
	}

	results <- Response{
		URL:        url,
		StatusCode: resp.StatusCode,
		Body:       string(body[:min(200, len(body))]),
		Duration:   time.Since(start),
	}
}

func main() {
	urls := []string{
		"https://google.com",
	}

	var wg sync.WaitGroup
	results := make(chan Response, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("Результаты")

	for result := range results {
		if result.Error != nil {
			fmt.Printf("\n❌ URL: %s\n", result.URL)
			fmt.Printf("   Ошибка: %v\n", result.Error)
			fmt.Printf("   Время: %v\n", result.Duration)
		} else {
			fmt.Printf("\n✅ URL: %s\n", result.URL)
			fmt.Printf("   Статус: %d\n", result.StatusCode)
			fmt.Printf("   Время: %v\n", result.Duration)
			fmt.Printf("   Начало ответа: %s...\n", result.Body)
		}
	}
}
