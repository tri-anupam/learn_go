package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	//interfaces
	eb := englishBot{}
	sb := spanishBot{}

	printGreetng(eb)
	printGreetng(sb)
}

func printGreetng(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	//very custom logic for generating an english greeting
	return "hi there!"
}

func (sb spanishBot) getGreeting() string {
	return "hola!"
}
