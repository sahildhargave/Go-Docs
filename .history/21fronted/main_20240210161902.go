//😎😍😘🥰😘😍😎

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Welcome to web verb video - https://google.com")

	PerformGetRequest()

}

func PerformGetRequest() {
	const myUrl = "http://localhost:3001/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

//This function is used for sending a post request

func PerformPostJsonRequest(){
	const myurl = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
	   {
		"coursename"  : "Let's go with golang",
		"price" : 0,
		"platform":"gogle.in"
	   }
	`)

	responsehttp.Post( myurl, "application/json", requestBody)
}