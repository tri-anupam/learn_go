package main

func main() {
	cards := deck{"Ace of Dimonds", newCard()} //here deck === []string{"Ace of Dimonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print() //here we call the receiver function

}

func newCard() string {
	return "Five of Dimonds"
}
