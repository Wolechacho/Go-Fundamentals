package main

import "fmt"

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Mark",
		age:  dog.Years(10),
	}

	fmt.Println(fido)
}
