package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition - Youtube")

	wg := &sync.WaitGroup{}

	var score = []int{0}
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		fmt.Println("One Routine")
		score = append(score, 1)
		wg.Done()
	}(wg)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Two Routine")
		score = append(score, 2)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Two Routine")
		score = append(score, 3)
		wg.Done()
	}(wg)
	wg.Wait()

}
