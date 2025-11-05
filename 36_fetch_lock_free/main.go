package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

type User struct {
	Name string
}

func main() {
	t := time.Now()
	names, err := Work(context.Background(), []User{
		{"Vasya"}, {"Petya"}, {"Petya"},
	})

	if err != nil {
		fmt.Println("Error in context: ", err)
		os.Exit(1)
	}

	fmt.Println("Result names: ", names)
	fmt.Println("Elapsed time: ", time.Since(t))
}

func networkWork(_ context.Context, u User) (string, error) {
	time.Sleep(time.Millisecond * 1)
	return u.Name, nil
}

func Work(ctx context.Context, users []User) (map[string]int64, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	names := make(map[string]int64)
	var firstError error

	for _, u := range users {
		user := u
		wg.Add(1)

		go func() {
			defer wg.Done()

			if ctx.Err() != nil {
				return
			}

			name, err := networkWork(ctx, user)

			mutex.Lock()
			defer mutex.Unlock()

			if firstError != nil {
				return
			}

			if err != nil {
				firstError = err
				cancel()
				return
			}

			names[name]++
		}()
	}

	wg.Wait()

	if firstError != nil {
		return nil, firstError
	}

	return names, nil
}
