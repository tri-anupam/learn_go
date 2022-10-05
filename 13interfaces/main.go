package main

type englishBot struct{}
type spanishBot struct{}

func main() {
	//interfaces

}

func (eb englishBot) getGreeting() string {
	//very custom logic for generating an english greeting
	return "hi there!"
}

func (sb spanishBot) getGreeting() string {
	return "hola!"
}
