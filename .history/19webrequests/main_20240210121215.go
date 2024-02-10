package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("Web Request")

	http.Get(url)

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response of type: %T", response)

	defer response.Body.Close() // caller' Responsility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}
