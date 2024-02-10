package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 23

	var result string
	if loginCount < 10 {
		result = "Regular user"
	}

	fmt.Println("Result", result)
}
