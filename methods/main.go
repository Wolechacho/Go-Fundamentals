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
	fmt.Println("I am", s.first, s.last)
}

func main() {
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

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	// p1 := person{
	// 	"Mark",
	// 	"Anderson",
	// }
	// p1.speak()

}
