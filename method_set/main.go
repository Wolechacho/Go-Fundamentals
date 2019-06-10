package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {

	//Method sets detemine what methods attach  to a TYPE.
	//The method set of a type determines the interface that the type implements
	//and the can be caled using a receiver of that type

	c := circle{5}
	// info(c)
	fmt.Println(c.area())
}
