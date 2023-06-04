package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(data chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		data <- i
	}
	close(data)
}

func consumer(data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range data {
		fmt.Println("Received:", num)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	// The unbuffered channel for communication
	data := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)

	// producer : produce the data and send to the channel
	go producer(data, &wg)
	// consumer : consume the data from the channel
	go consumer(data, &wg)

	// Wait for all goroutines to finish
	wg.Wait()
}
