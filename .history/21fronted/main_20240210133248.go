//😎😍😘🥰😘😍😎

package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to web verb video - https://google.com")

	Pe

}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
}

//This function is used for sending a post request
