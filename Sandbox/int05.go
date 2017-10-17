package main

import (
	"fmt"
	"math"
)

//Shaper is an interface and has a single function Area that returns a float
type Shaper interface {
   Area() float64
   Perimeter() float64
}

type Rectangle struct {
   length, width int
}

type Circle struct {
	radius int
}

func (r Rectangle) Area() float64 {
   return float64(r.length * r.width)
}

func (r Rectangle) Perimeter() float64 {
   return float64( 2 * (r.length + r.width))
}

func (c Circle) Area() float64 {
	return float64(c.radius * c.radius) * math.Pi
}

func (c Circle) Perimeter() float64 {
	return float64(c.radius) * math.Pi * 2
}

func showShape(s Shaper) {
	fmt.Printf("The shape's area %v\n", s.Area())
	fmt.Printf("The shape's perimeter %v\n", s.Perimeter())
}

func main() {
   r := Rectangle{length:5, width:3}
   c := Circle{radius:7}

   showShape(r)
   showShape(c)
}
