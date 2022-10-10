package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Kiwi"
	fruitList[2] = "Pineapple"
	fruitList[3] = "Mango"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Length of fruitlist is ", len(fruitList))

	var vegList = [3]string{"potatos", "tomatos", "carrot"}
	fmt.Println("Very list is ", vegList)

}
