package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Если произошла ошибка при закрытии resp.Body.Close(), надо ее обработать и сделать это можно несколькими способами
// Если это не сделать, может быть утечка соединений и памяти

func fetchURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("запрос не удался: %w", err)
	}

	defer func() {
		if resp.Body != nil {
			// Вычитываем остатки чтобы Close с большей вероятностью прошел
			_, _ = io.Copy(io.Discard, resp.Body)

			if closeErr := resp.Body.Close(); closeErr != nil {
				log.Printf("Body Close error: %v", closeErr)
				resp.Close = true
			}
		}
	}()

	// Читаем Body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("чтение не удалось: %w", err)
	}

	return body, nil
}

func main() {
	body, err := fetchURL("https://google.com")

	if err != nil {
		fmt.Println("Ошибка чтения содержимого сайта")
	}

	fmt.Println(body)
}
