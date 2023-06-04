package main

import (
	"fmt"
	"time"
)

func producer(data chan<- int) {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Send:", i)
		data <- i
	}
	close(data)
}

func consumer(data <-chan int, done chan<- bool) {
	for num := range data {
		fmt.Println("Received:", num)
	}
	done <- true
}

func main() {
	// The channels for communication
	data := make(chan int)
	done := make(chan bool)

	// producer : produce the data and send to the channel
	go producer(data)

	// consumer : consume the data from the channel
	go consumer(data, done)

	// Wait for the consumer to finish
	<-done
}
