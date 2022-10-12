package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://tri-anupam.github.io/gaming-website.github.io/ z"

func main() {
	fmt.Println("web Requests")

	response, err := http.Get(url)
	errNilChecking(err)
	// fmt.Println(response)
	fmt.Printf("\n\nthe type of response is %T", response)

	defer response.Body.Close() //caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)
	errNilChecking(err)
	content := string(databytes)
	fmt.Println(content)

}

func errNilChecking(err error) {
	if err != nil {
		panic(err)
	}
}
