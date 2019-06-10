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

type person2 struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
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
	fmt.Println("This is the Marshed byte: ", string(bs))

	s := `[{"First":"Mike","Last":"Segun","Age":54},{"First":"Ade","Last":"Sally","Age":39}]`
	fmt.Printf("%T\n", s)
	bsNew := []byte(s)
	fmt.Printf("%T\n", bsNew)
	fmt.Println(bsNew)
	fmt.Println()

	var people2 []person2
	// people2 := []person2{}

	err2 := json.Unmarshal(bsNew, &people2)

	if err != nil {
		fmt.Println("This is the err: ", err2)
	}

	fmt.Println()
	fmt.Println("All the data", people2)

	for i, v := range people2 {
		fmt.Println("\nPerson Number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
