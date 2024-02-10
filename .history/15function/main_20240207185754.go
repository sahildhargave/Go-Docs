package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
}

func greeter() {
	fmt.Println("Namstey from golang")
}

// in golang func inside func cannot be called or written
