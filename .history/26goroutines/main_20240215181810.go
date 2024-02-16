package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//go greeter("Hello")
	//greeter("world")
	

}

//func greeter(s string) {
//	for i := 0; i < 6; i++ {
//		time.Sleep(3 * time.Millisecond)
//		fmt.Println(s)
//	}
//}

//func greateNight(s string) {
//	for i := 0; i < 6; i++ {
//		time.Sleep(3 * time.Millisecond)
//		fmt.Println("Sahil Dhargave")
//	}
websitelist := []string{
	"https://youtube.com",
	"http://go.dev",
	"https://google.com",
	"https://github.com",
	"https://chatgtp.com",
}

for _, web := range websitelist{
	getStatusCode(web)
}
}

func getStatusCode(endpoint string){
	res,err := http.Get(endpoint)
     
	if err!=nil {
		fmt.Println("OOPS in endpoint")
	}
	fmt.Println("200 status code for website is %s",res.StatusCode, endpoint)

}