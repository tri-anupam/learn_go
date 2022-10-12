package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
	loop()
}

func loop() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
