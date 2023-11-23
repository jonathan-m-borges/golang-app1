package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	messages := make(chan string, 5)

	go consumer(messages, "a")
	go consumer(messages, "b")
	go producer(messages, "c")
	go producer(messages, "d")

	time.Sleep(time.Second * 50)
}

func consumer(messages chan string, label string) {
	for x := range messages {
		fmt.Println("consumer", label, x, "len", len(messages))
		time.Sleep(time.Millisecond * 200)
	}
}

func producer(messages chan string, label string) {
	for i := 0; i < 10; i++ {
		messages <- strconv.Itoa(i)
		fmt.Println("producer", label, i, "len", len(messages))
		time.Sleep(time.Millisecond * 50)
	}
}
