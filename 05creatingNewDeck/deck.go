package main

import "fmt"

type deck []string

//Here we are creating a new deck--------->>>>>>>
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four" /*"Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jocker", "Queen", "King"*/}

	for _, suit := range cardSuits { ///'_' underscore is used a variable for a variable
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit) //append is used because cards slice string is empty
		}
	}
	return cards
}

func (d deck) print() { //'d' is the variable of type deck
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
