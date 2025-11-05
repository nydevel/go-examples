package main

import (
	"fmt"
	"runtime"
)

// Количество логических ядер
// если запустить программу с GODEBUG=schedtrace=1000 ./main то отобразится отладочная информация о процессах и тп

func main() {
	fmt.Println(runtime.NumCPU())

	// Количество процессоров (GMP - P) - с нулевым параметром выводит количество процессоров, если больше - устанавливает количество
	fmt.Println(runtime.GOMAXPROCS(0))
}
