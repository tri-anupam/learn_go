package main

import "fmt"

var num int = 10
var num2 int

func main() {
	// var cards string = "Ace of Spades"

	cards := "Ace of Spades" //This is another type to initlize a variable in golang

	cards = "Five of Dimonds" //to update the value of variable use only "=" sign

	num1 := 1 //Assigning a number variable

	var decksize int //We can declare a variable and then assign it later
	decksize = 52
	fmt.Println(decksize)

	fmt.Println(cards + cards)
	fmt.Println(num1)

	fmt.Println(num)
	num2 = 20
	fmt.Println(num2)
	fmt.Println(num + num2)

}
