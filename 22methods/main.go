package main

import "fmt"

func main() {
	fmt.Println("Methods in golang ")

	anupam := User{"Anupam", "anupam@gmail.com", true, 13}
	fmt.Println(anupam)
	anupam.GetStatus()
	anupam.NewMail()
	// fmt.Println(anupam)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() { //methods are receiver function in golang
	fmt.Println(u.Status)
}

func (u User) NewMail() {
	u.Email = "anupamtri@gmai.com"
	fmt.Println(u.Email)
}
