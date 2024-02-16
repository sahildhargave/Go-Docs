package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Docs to the Channel ")

	myCh := make(chan int, 5)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	//fmt.Println(<-myCh)
	//myCh <- 5

	go func(ch chan int, wg *sync.WaitGroup) {
		val, isChanelOpen := <-myCh
		fmt.Println(isChanelOpen)
		fmt.Println(val)
		fmt.Println(<-myCh)
		//  if we dont want to do any operation for
		// channel then add buffer in the channer
		// as above myCh := make(chan int,2)
		// here 2 is buffer
		// here 1 is pass then its a deffault one
		//fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		myCh <- 0
		myCh <- 5
		myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
