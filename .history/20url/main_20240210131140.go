package main

import (
	"fmt"
	"net/url"
)

const website string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Learning to handling url in golang")
	fmt.Println(website)

	result, _ := url.Parse(website)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Hostname())
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are : %T\n", qparams)
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for _, value := range qparams {
		fmt.Println("Param is:", value)
	}

	partsOfUrl := &url.URL{
       Scheme: "https",
	   Host: "loc.dev",
	   Path: "/tutcss",
	   RawPath: "user=sahil",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

	fmt.Println(partsOfUrl.Scheme)
}
