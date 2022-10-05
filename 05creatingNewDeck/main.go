package main

func main() {
	cards := newDeck() //here cards is the variable that call the function newDeck()

	//cards.print() //here print() is the receiver function that is used to call the all the values that is stored in cards where cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
