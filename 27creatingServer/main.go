package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to GOooooo")

	//PerformGetRequest()            //to perform get request

	// PerformPostRequest()          //to perform post request

	PerformPostFormRequest() //to perform post form request

}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json paylaod

	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price":0,
			"platform":"www.google.com"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "Anupam")
	data.Add("lastname", "Tripathi")
	data.Add("email", "askdf@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll((response.Body))
	fmt.Println(string(content))
}
