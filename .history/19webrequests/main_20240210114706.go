package main

import (
	"fmt"
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
	fmt.Println("Response of type: %T", response)

	response.Body.Close() // caller' Responsility to close the connection

}
