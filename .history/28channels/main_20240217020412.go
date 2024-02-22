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
    
	// below should add ch <- chan int argument for clean code
	// recieve only
	// if it is recieve only then you are not allowed to 
	// close(myCh)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// close(myCh)
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
    
	//Send only
	go func(ch chan int, wg *sync.WaitGroup) {
		myCh <- 0
		myCh <- 5
		myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
