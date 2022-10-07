package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	side float64
}

type shape interface {
	getArea() float64
}

func main() {
	areaOfTriangle := triangle{2, 3}
	areaOfSquare := square{3}

	printArea(areaOfTriangle)
	printArea(areaOfSquare)
}

func (t triangle) getArea() float64 {
	return t.height * t.base * 0.5
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func printArea(s shape) {
	fmt.Println(s.getArea())

}
