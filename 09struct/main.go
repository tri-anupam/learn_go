package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo contactInfo
}

func main() {
	// anupam := person{"Anupam", "Trpathi"}
	// fmt.Println(anupam)

	// anubhav := person{firstName: "Anubhav", lastName: "Singh"}
	// fmt.Println(anubhav)

	//👆👆Updating Struct Values👆👆
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	//embedding Structs 🔱🔱🔱
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}
	fmt.Println(jim)

}
