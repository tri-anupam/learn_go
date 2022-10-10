package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointer")

	// var ptr *int                         //initlize a pointer
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber //reference means '&'

	fmt.Println("Value of actual pointer is ", ptr)  //it store the address of variable myNumber
	fmt.Println("Value of actual pointer is ", *ptr) // it stores the value of variable myNumber
	fmt.Println("Value of actual pointer is ", &ptr) // it stores the address of pointer ptr

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)
}
