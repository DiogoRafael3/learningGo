package main

import (
	"fmt"
	"math"
)

type shape interface { //to create an interface, sort of like a struct: type <NAME> interface
	area() float64 //this means every shape should have an area function
}

type rectangle struct {
	height float64
	width  float64
}

func (r rectangle) area() float64 { //by implementing area(), rectangle is now implicitly a shape thanks to the interface above defined
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 { //different implementation of the same method
	return math.Pi * math.Pow(c.radius, 2)
}

func printArea(s shape) {
	fmt.Println("Area is:", s.area())
}

func main() {
	r := rectangle{10, 15}
	printArea(r)

	c := circle{10}
	printArea(c)
}
