//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//func main() {
//	fmt.Println("Docs to the Channel ")
//
//	myCh := make(chan int, 2)
//	wg := &sync.WaitGroup{}
//	wg.Add(2)
//	//fmt.Println(<-myCh)
//	//myCh <- 5
//    
//	// below should add ch <- chan int argument for clean code
//	// recieve only
//	// if it is recieve only then you are not allowed to 
//	// close(myCh)
//	go func(ch <-chan int, wg *sync.WaitGroup) {
//		// close(myCh)
//		val, isChanelOpen := <-myCh
//		fmt.Println(isChanelOpen)
//		fmt.Println(val)
//		fmt.Println(<-myCh)
//		//  if we dont want to do any operation for
//		// channel then add buffer in the channer
//		// as above myCh := make(chan int,2)
//		// here 2 is buffer
//		// here 1 is pass then its a deffault one
//		//fmt.Println(<-myCh)
//		wg.Done()
//	}(myCh, wg)
//    
//	//Send only
//	go func(ch chan int, wg *sync.WaitGroup) {
//		myCh <- 0
//		myCh <- 5
//		myCh <- 6
//		close(myCh)
//		wg.Done()
//	}(myCh, wg)
//	wg.Wait()
//}

package main

import( 
	"net/http"
	"fmt"
)
func main(){
 links := []string{
	"http://google.com",
	"http://facebook.com",
	"http://stackoverflow.com",
	"http://youtube.com",
 }

 ch := make(chan string)

 for _,link := range links{
	//fmt.Println("link:",link)
      go checkLink(link, ch)

 }
 fmt.Println(<- ch)
}

func checkLink(link string,ch chan string){
	_,err :=http.Get(link)
	if err!=nil{
		fmt.Println(link,"might be down!")
		ch <- "Might be down I think"
		return
	}
	fmt.Println(link,"is up!")
	ch <- "Yep its up"
}