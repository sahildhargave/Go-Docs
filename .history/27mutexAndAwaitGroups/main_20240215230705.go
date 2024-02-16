package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition - Youtube")

	wg := &sync.WaitGroup{}

	var score = []int{0}
	go func() {
		fmt.Println("One Routine")
		score = append(score, 1)
	}()
	go func() {
		fmt.Println("Two Routine")
		score = append(score, 2)
	}()
	go func() {

	}()

}
