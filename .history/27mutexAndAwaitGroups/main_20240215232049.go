package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition - Youtube")

	wg := &sync.WaitGroup{}
	var mu sync.Mutex

	var score = []int{0}
	wg.Add(3)

	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()
		fmt.Println("One Routine")
		mu.Lock()
		score = append(score, 1)
		mu.Unlock()
	}(wg, &mu)

	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Two Routine")
		mu.Lock()
		score = append(score, 2)
		mu.Unlock()
	}(wg, &mu)

	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Three Routine")
		mu.Lock()
		score = append(score, 3)
		mu.Unlock()
	}(wg, &mu)

	wg.Wait()
	fmt.Println(score)
}
