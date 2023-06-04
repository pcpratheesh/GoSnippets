package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS

	switch os {
	case "linux":
		fmt.Println("Running on Linux")
	case "darwin":
		fmt.Println("Running on macOS")
	case "windows":
		fmt.Println("Running on Windows")
	default:
		fmt.Println("Unknown operating system")
	}
}
