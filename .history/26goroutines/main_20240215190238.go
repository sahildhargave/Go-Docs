package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup //pointer

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
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		signals = append(signals, endpoint)

		fmt.Printf("200 status code for website is %s\n", res.StatusCode, endpoint)
	}

}
