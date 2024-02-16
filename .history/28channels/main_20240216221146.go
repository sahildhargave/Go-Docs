package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Docs to the Channel ")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	//fmt.Println(<-myCh)
	//myCh <- 5

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		myCh <- 5
		myCh <- 6
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
