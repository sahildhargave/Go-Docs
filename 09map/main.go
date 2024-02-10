package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in go")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["go"] = "Golang"

	fmt.Println("List of all languages:", languages)
	fmt.Println("Js short for", languages["js"])
	delete(languages, "rb")
	fmt.Println("List of all languages:", languages)

	// control and loop are interesting in golang

	for key, value := range languages {
		fmt.Printf("For key %v , value id %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("For key %v , value id %v\n", value)
	}
}
