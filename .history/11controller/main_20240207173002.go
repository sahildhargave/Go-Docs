package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 34

	var result string
	if loginCount < 10 {
		result = "Regular user"
	} else {
		result = "something else"
	}

	fmt.Println("Result", result)
}
