package main

import "fmt"

func main() {
	object := map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
	}
	// fmt.Println(object[3])
	// fmt.Println(object[2])
	// object[1] = "NoOne" //to update in map

	object[4] = "Four" //to add new key:value pair
	// fmt.Println(object)

	// delete(object, 1) //to delete a key value pair from map

	// val1 := object[2]
	// fmt.Println(val1) //retrive a value

	readMap(object)
}

func readMap(c map[int]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
