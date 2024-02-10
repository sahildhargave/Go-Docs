package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Ludo game!")

	rand.Seed(time.Now().UnixNano())

}
