package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

//any value of the type "secretAgent", have access to the speak method
//ie, the method speak will only be called by the receiver "secretAgent"
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

//keyword identifier type
type human interface {
	//the interface says, if u have the method below, then u are my type
	speak()
}

//demonstrating polymorphism
func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar new", h.(person).first)

	case secretAgent:
		fmt.Println("I was passed into bar new", h.(secretAgent).first)
	}
	// fmt.Println("I was passed into bar", h)
}

func main() {
	//A value can be of more than one type
	// e.g, "sa1" has type "secretAgent" and "human"
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Ali",
			"Mohammed",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

}
