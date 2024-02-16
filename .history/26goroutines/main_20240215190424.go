package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup

func main() {
	websitelist := []string{
		"https://youtube.com",
		"http://go.dev",
		"https://google.com",
		"https://github.com",
		"https://chatgpt.com",
	}

	for _, web := range websitelist {
		wg.Add(1)
		go getStatusCode(web)
	}

	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Printf("Error in endpoint %s: %v\n", endpoint, err)
	} else {
		signals = append(signals, endpoint)
		fmt.Printf("Status code for %s is %d\n", endpoint, res.StatusCode)
	}
}
