package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Docs to the Channel ")

	myCh := make(chan int)
	wg := &sync.WaitGroup{}
   
	//fmt.Println(<-myCh)
	//myCh <- 5

	go func(ch chan int, wg *sync.WaitGroup) {
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		wg.Done()
	}(myCh, wg)

}
