//😎😍😘🥰😘😍😎

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - https://google.com")

	PerformGetRequest()
	PerformPostJsonRequest()

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

func PerformPostJsonRequest() {
	const myurl = "http://localhost:3001/post"

	//fake json payload

	requestBody := strings.NewReader(`
	   {
		"coursename"  : "Let's go with golang",
		"price" : 0,
		"platform":"gogle.in"
	   }
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// formdata

	data := url.Values{}
	data.Add("first name", "Sahil")
	data.Add("Last name", "Dhargave")
	data.Add("Email", "sahildhargave5253@gmail.com")
	data.Add("Mobile No", "7756811710")
}
