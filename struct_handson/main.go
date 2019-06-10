package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.length * s.length
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func calArea(s shape) {
	// fmt.Println(s.area())
	switch s.(type) {
	case square:
		fmt.Println("This is the area of the square: ", s.area())
	case circle:
		fmt.Println("This is the area of the circle: ", s.area())
	}
}

func main() {
	c := circle{
		radius: 5.34,
	}
	s := square{
		length: 20.4,
	}

	calArea(c)
	calArea(s)
}
