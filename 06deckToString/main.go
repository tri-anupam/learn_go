package main

func main() {
	// cards := newDeck()
	// //fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.print()
}
