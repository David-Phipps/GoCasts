// Write a program that creates two custom sturct types called triangle and square. Use them in an interface with getArea function
package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	t := triangle{
		height: 1.2,
		base:   4.5,
	}

	s := square{
		sideLength: 5.0,
	}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 { //we put triangle receiver first, will return float64
	area := 0.0
	area = 0.5 * t.height * t.base
	return area
}

func (s square) getArea() float64 { // we put square receiver first, will return float64G
	square := 0.0
	square = s.sideLength * s.sideLength
	return square
}
