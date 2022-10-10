package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now() //initlize the date
	fmt.Println(presentTime.Date())

	fmt.Println(presentTime.Format("01-02-2006"))
}
