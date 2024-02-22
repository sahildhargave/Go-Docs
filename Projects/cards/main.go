package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome")
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
