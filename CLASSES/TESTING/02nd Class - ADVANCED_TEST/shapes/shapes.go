package shapes

import (
	"fmt"
	"math"
)

// interface created in class 16 of the basic classes, to use this in other files we must upper case every function and field
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func PrintArea(s Shape) {
	fmt.Println("Area is:", s.Area())
}
