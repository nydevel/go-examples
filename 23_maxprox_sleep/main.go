package main

import (
	"fmt"
	"runtime"
	"time"
)

// При GOMAXPROCS=1 у нас только один поток ОС, который выполняет код Go.
// time.Sleep() - это блокирующий вызов, но он не занимает CPU
// Когда горутина вызывает Sleep(), планировщик Go переключается на другую горутину
// Все 10 горутин "спят" параллельно (даже на одном потоке!)
func main() {
	runtime.GOMAXPROCS(1) // Один поток!

	start := time.Now()

	for i := 0; i < 10; i++ {
		go func(id int) {
			time.Sleep(1 * time.Second)
			fmt.Println("Горутина", id, "завершена")
		}(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Время выполнения:", time.Since(start))
}
