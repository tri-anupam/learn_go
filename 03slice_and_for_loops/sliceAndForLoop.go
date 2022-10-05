package main

import "fmt"

func main() {
	//Creating an array
	// arr := [7]string{"one", "two", "three", "four", "five", "six", "seven"}

	//Display array
	// fmt.Println("Array: ", arr)

	//Creating a slice
	// mySlice := arr[1:6]

	//Display the slice
	// fmt.Println("Slice: ", mySlice)

	//Display length of the slice
	// fmt.Println("Length of the Slice is: ", len(mySlice))

	//Display the capacity of the slice
	// fmt.Println("Capacity of the slice is: ", cap(mySlice))

	// fmt.Println("Capacity of the array is: ", cap(arr))
	// fmt.Println("Length of the array is: ", len(arr))

	//Initlization of slice 🔪🔪
	cards := []string{"Ace of Dimonds", newCard()}

	//To add new element in slice
	cards = append(cards, "Six of Spades")

	// fmt.Println(cards)

	//Iterating over a slice of cards
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
func newCard() string {
	return "Five of Dimonds"
}
