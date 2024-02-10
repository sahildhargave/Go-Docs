package main 

import "fmt"

const url = "https://google.com";

func main(){
	fmt.Println("Web Request");
    
	http.Get(url)

	response, err := http.Get(url)

	if err != nil{
		panic(err)
	}
	fmt.Println("Response of type: %T",response)

	
}