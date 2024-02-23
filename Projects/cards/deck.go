package main

import (
	"fmt"
    "strings"
)

type deck []string

func newDeck() deck{
cards := deck{}
cardSuit := []string{"Spades","Diamonds","Hearts","Clubs"}
cardValues := []string{"Ace","Two","Three","Four"}
 for i,suit := range cardSuit{
	 for j,value := range cardValues{
		 cards = append(cards,value+" of "+suit)
		 fmt.Println(i,j)
	 }
 }
 return cards
}
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// tell the function to return two value of deak
//return two separate value
func deal(d deck ,handSize int)(deck, deck){
 return d[:handSize],d[handSize:]
}

func (d deck) toString() string{
	return strings.Join([]string(d),",")
	
}
