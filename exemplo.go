package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan int)

	go produtor(canal, "P1")
	go produtor(canal, "P2")

	go consumidor(canal, "C1")
	go consumidor(canal, "C2")
	consumidor(canal, "C3")
}

func produtor(canal chan int, label string) {
	for i := 0; i < 10; i++ {
		canal <- i
		fmt.Println("produtor", label, "produziu", i)
	}
}

func consumidor(canal chan int, label string) {
	for x := range canal {
		fmt.Println("consumidor", label, "consumiu", x)
		time.Sleep(time.Second)
	}
}
