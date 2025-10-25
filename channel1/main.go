package main

import "fmt"

func printMessage(messageChannel chan string) {
	for message := range messageChannel {
		fmt.Println(message)
	}
}

func processMessages(messages []string) {
	messagesCh := make(chan string)
	go printMessage(messagesCh)

	for _, message := range messages {
		messagesCh <- message
	}

	defer close(messagesCh)
}

func main() {
	messages := []string{"message1", "message2"}
	processMessages(messages)
}
