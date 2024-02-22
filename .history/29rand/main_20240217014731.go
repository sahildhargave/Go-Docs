package main

import (
	"crypto/rand"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//incorrect
	//	var One int = 2
	//	var Two float64 = 4
	//
	//	fmt.Println("The sum is:", One+Two)
	//random number
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(5) + 1)
}
