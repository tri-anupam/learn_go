package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo //We can write it in this case
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Carter",
		contactInfo: contactInfo{
			email:   "jim234@gmail.com",
			zipCode: 158798,
		},
	}

	//jimPointer := &jim  //jimPointer stores the address of jim
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {
	fmt.Println("%+v", p)
}
