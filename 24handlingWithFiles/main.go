package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcomes to files in golang")
	content := "Hello my name is Anupam Tripathi. Thank You!"

	file, err := os.Create("./myname.txt") //Creating a file

	checkNilErr(err)

	length, err := io.WriteString(file, content) //writing in the file with the content

	checkNilErr(err)

	fmt.Println("Length of file is ", length)
	defer file.Close() //Close the file

	readFile("./myname.txt")
}

func readFile(filename string) { //to read the data from the file
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Print("text data inside the file is \n", string(databyte))
}

func checkNilErr(err error) { //this is error handling fiel
	if err != nil {
		panic(err)
	}
}
