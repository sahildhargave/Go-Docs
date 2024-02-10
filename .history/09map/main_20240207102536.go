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
}
