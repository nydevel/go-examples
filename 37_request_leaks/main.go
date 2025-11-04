package main

import (
	"log"
	"net/http"
)

// Без defer resp.Body.Close() программа зависнет
// после ~100 запросов

func main() {
	for i := 0; i < 1000; i++ {
		resp, err := http.Get("https://google.com")
		if err != nil {
			return // Body нет, закрывать нечего
		}

		// !!! defer resp.Body.Close() обязательно нужен чтобы не было утечек памяти и чтобы не было переполнения пула запросов, через defer чтобы выполнился несмотря на панику
		defer func() {
			// если перед этим не было попытки выполнить io.ReadAll и обработки ее ошибки, то надо обрабатывать ошибку в resp.Body.Close()
			if err := resp.Body.Close(); err != nil {
				log.Printf("Предупреждение: ошибка при закрытии Body: %v", err)
			}
		}()
	}
}
