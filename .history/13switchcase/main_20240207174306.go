package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Ludo game!")

	rand.Seed(time.Now().UnixNano())
	// Create a new game with 4 players and 12 dice per player engine
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)

}
