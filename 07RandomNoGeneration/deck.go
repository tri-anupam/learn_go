package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

// Here we are creating a new deck--------->>>>>>>
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four" /*"Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jocker", "Queen", "King"*/}

	for _, suit := range cardSuits { //'_' underscore is used a variable for a variable
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
func deal(d deck, handSize int) (deck, deck) { //new function is created with multiple return values
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",") //Join concatenates the elements of its first argument to create a single string. The separator string sep is placed between elements in the resulting string.
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) //bs=byteSlice and err = value of type 'error', if nothing went wrong it will have a value of 'nil'
	if err != nil {
		//Option #1 - log the error and return a call to newDeck()
		//Option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() { //Random number generation
	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
