package main

import (
	"fmt"
	"math"
)

type person struct {
	name string
}

func changeMe(p *person) {
	// (*p).name = "Ali Judge"
	p.name = "Ali Judge"

}

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
	fmt.Println("area: ", s.area())
}

func main() {

	c := circle{6}
	info(&c)

	p1 := person{
		name: "Mark Essien",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)

	a := 23
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	//when you have the address and you wan the value stored in the address
	// & gives us the address, * gives us the values of the address
	fmt.Println(*b)
	fmt.Println(*&a)

	fmt.Println()

	x := 0
	fmt.Println("x address before ", &x)
	fmt.Println("x before", x)
	foo(&x)
	fmt.Println("x address after", &x)
	fmt.Println("x after: ", x)

}

func foo(y *int) {
	fmt.Println("y before", y)
	fmt.Println("y value for *: ", *y)
	//dereferencing:
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after with * ", *y)

}
