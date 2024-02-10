//😎😍😘🥰😘😍😎

package main

import (
	"fmt"

	"golang.org/x/tools/go/analysis/passes/defers"
)

func main(){
	fmt.Println("Welcome to web verb video - https://google.com")

}

func PerformGetRequest(){
  const myUrl = "http://localhost:8000/get"

  response, err := http.GET(myUrl)

  if err != nil {
	panic(err)
  }

  defer response.Body.Close();

  fmt.Println("Status code: ")
}



//This function is used for sending a post request

