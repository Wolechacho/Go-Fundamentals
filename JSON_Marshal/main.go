package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Mike",
		Last:  "Segun",
		Age:   54,
	}

	p2 := person{
		First: "Ade",
		Last:  "Sally",
		Age:   39,
	}

	people := []person{p1, p2}

	fmt.Println(people)
	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
