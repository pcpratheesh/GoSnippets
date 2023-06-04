package main

import "fmt"

func main() {
	defer fmt.Println("this will run at the end")
	defer fmt.Println("This will run right before the end defer")
}
