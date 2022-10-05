package main

import "fmt"

//Create a new type of 'deck' which is slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() { //here print() is the receiver fucntion which are created by us
	for i, card := range d {
		fmt.Println(i, card)
	}
}
