package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now()
	names, err := Do(context.Background(), []User{
		{"Paul"}, {"Jack"}, {"Jack"}, {"Mike"},
	})
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println("Result: ", names)
	fmt.Println("Elapsed time: ", time.Since(t))
}

// fetch что-то дeлaeм по сeти.
func fetch(_ context.Context, u User) (string, error) {
	time.Sleep(time.Millisecond * 10) // имитaция зaдeржки
	return u.Name, nil
}

func Do(ctx context.Context, users []User) (map[string]int64, error) {
	names := make(map[string]int64, len(users))

	for _, u := range users {
		name, err := fetch(ctx, u)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		names[name] = names[name] + 1
	}

	return names, nil
}
