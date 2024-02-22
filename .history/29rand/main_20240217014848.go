package main

import (
	"crypto/rand"
)

func main() {
	//incorrect
	//	var One int = 2
	//	var Two float64 = 4
	//
	//	fmt.Println("The sum is:", One+Two)
	//random number
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(5) + 1)
	// random with crypto
	randNumber, _ := rand.Int(rand.Reader)
}
