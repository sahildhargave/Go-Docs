package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(34, 35)

	fmt.Println("Addition of two is", result)
}

func adder(x int, y int) int {
	return x + y
}
func greeter() {
	fmt.Println("Namstey from golang")
}

// in golang func inside func cannot be called or written
