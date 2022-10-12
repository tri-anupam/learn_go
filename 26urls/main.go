package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://www.google.com/search?q=hello+world&rlz=1C1RXQR_enIN1025IN1025&oq=hello+world&aqs=chrome.0.35i39j0i433i512j0i131i433i512j0i512l3j0i131i433i512j46i433i512j0i512j0i131i433i512.3088j0j15&sourceid=chrome&ie=UTF-8"

func main() {
	fmt.Println("Welcome to handling url in golang")
	fmt.Println(myUrl)

	//parsing
	result, _ := url.Parse(myUrl)
	//fmt.Println(result.Scheme)
	//fmt.Println(result.Host)
	//fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	//fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}
}
