package main

import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)

}

func newCard() string { //when assigning a function in golang we must declare the datatype of the function
	return "Five of Dimonds" //it returns only the type of data that is metion
}
