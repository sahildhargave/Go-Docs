package main

import "fmt"

func main() {
	greeter("Hello")
	greeter("world")
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		fmt.Println(s)
	}
}
