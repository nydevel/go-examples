package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func fetchURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Ошибка при запросе %s: %v\n", url, err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Ошибка получения тела ответа: %v", err)
		return
	}

	fmt.Printf("Тело ответа: %s\n", body)
}

func main() {
	var wg sync.WaitGroup

	urls := []string{
		"https://google.com",
		"https://wildberries.ru",
	}

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, &wg)
	}

	wg.Wait()
	fmt.Println("Все запросы завершены!")
}
