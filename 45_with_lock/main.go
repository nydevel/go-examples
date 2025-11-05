package main

import (
	"fmt"
	"sync"
)

func withLock(mutex *sync.Mutex, action func()) {
	if action == nil {
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	action()
}

func doSomething() error {
	return nil
}

func main() {
	mutex := sync.Mutex{}

	withLock(&mutex, func() {
		err := doSomething()

		if err == nil {
			fmt.Println("changed")
		}
	})
}
