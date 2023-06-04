package main

import (
	"errors"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
func main() {
	_, err := Divide(10, 0)
	if err != nil {
		panic(err)
	}
}
