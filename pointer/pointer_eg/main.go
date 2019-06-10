package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello")
}

func sayHello(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Mark",
	}
	sayHello(&p1)
}
