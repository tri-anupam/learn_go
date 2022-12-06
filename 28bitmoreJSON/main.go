package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string
}

func main() {
	fmt.Println("Welcome to JSON")
	EncodeJson()
}

func EncodeJson() {
	allCourses := []course{
		{"ReactJS Bootcamp", 299, "Youtube", "abc123", []string{"web-dev", "js"}},
		{"Python Bootcamp", 199, "Youtube", "absdc123", []string{"back-end", "py"}},
		{"Html Bootcamp", 299, "Youtube", "asf3123", nil},
	}

	//package this data as JSON data

	finalJson, err := json.Marshal(allCourses)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
