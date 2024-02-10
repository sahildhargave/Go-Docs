package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 10

	var result string
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "something else"
	} else {
		result = "😀"
	}

	fmt.Println("Result", result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}
}
