package main

import "fmt"

func main() {
	fmt.Println("Docs to the Channel ")

	myCh := make(chan int)

	myCh <- 5
	fmt.Println(<-myCh)
}
