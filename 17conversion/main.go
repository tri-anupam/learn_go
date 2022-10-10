package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin) // initlize to take user input

	input, _ := reader.ReadString('\n') //taking user input

	fmt.Println("Thanks for rating, ", input) //print the user input

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // convert it into float

	if err != nil { //handling with error
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1) //print the user input in float form from string
	}

}
