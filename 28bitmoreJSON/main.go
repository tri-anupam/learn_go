package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string   `json:"-"`              //'-' dash don't reflect the value to the user but it stores it.
	Tags     []string `json:"tags,omitempty"` //'omitempty don't reflect the field which is empty.'
}

func main() {
	fmt.Println("Welcome to JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	allCourses := []course{
		{"ReactJS Bootcamp", 299, "Youtube", "abc123", []string{"web-dev", "js"}},
		{"Python Bootcamp", 199, "Youtube", "absdc123", []string{"back-end", "py"}},
		{"Html Bootcamp", 299, "Youtube", "asf3123", nil},
	}

	//package this data as JSON data

	finalJson, err := json.MarshalIndent(allCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

//how to consume JSON data in golang

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"Platform": "Youtube",
		"tags": ["web-dev","js"]
	}
	`)

	var yourCourse course

	checkValid := json.Valid(jsonDataFromWeb) //give the boolean value

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &yourCourse) //'&' with yourCourse gave the reference to it.
		fmt.Printf("%#v\n", yourCourse)
	} else {
		fmt.Println("your json is not valid")
	}

	//some cases where you want to add to kay value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is: %T\n", k, v, v)
	}

}
