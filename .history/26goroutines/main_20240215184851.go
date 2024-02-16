package main

import (
	"fmt"
	"net/http"
)

func main() {

	websitelist := []string{
		"https://youtube.com",
		"http://go.dev",
		"https://google.com",
		"https://github.com",
		"https://chatgtp.com",
	}

	for _, web := range websitelist {
		getStatusCode(web)
	}
}

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	}
	fmt.Printf("200 status code for website is %s", res.StatusCode, endpoint)

}
