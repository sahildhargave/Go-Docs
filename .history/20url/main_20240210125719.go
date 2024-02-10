package main

import (
	"fmt"
	"net/url"
)

const website string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Learning to handling url in golang")
	fmt.Println(url)

	result, _ := url.Parse(website)

	fmt.Println(result.Scheme)

}
