package main

import "fmt"

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
