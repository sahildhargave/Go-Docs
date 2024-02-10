package main 

import "fmt"

const url = "https://google.com";

func main(){
	fmt.Println("Web Request");
    
	http.Get(url)
}