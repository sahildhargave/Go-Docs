package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 34

	var result string
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "something else"
	} else {
		result = "😀"
	}

	fmt.Println("Result", result)
}
