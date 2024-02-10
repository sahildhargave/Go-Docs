package main

import "fmt"

func main() {
	fmt.Println("Start")

	// This function call will be deferred until the surrounding function (main) completes.
	defer cleanup()

	fmt.Println("Processing...")

	// Other code here

	fmt.Println("End")
	myDefer()
}

func cleanup() {
	fmt.Println("Performing cleanup tasks")
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
