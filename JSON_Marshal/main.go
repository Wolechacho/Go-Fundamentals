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

type user struct {
	Name    string
	Age     int
	Sayings []string
}

type userNew struct {
	Name    string   `json:"Name"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	user1 := user{
		Name: "Steven",
		Age:  30,
		Sayings: []string{
			"You are awesome",
			"You are great",
		},
	}
	user2 := user{
		Name: "Kedu",
		Age:  40,
		Sayings: []string{
			"This is him",
			"He is an awesome teacher",
		},
	}
	users := []user{user1, user2}

	fmt.Println("This is the sum: ", users)

	buser, erruser := json.Marshal(users)
	if erruser != nil {
		fmt.Println("Error occured: ", erruser)
	}
	fmt.Println("This is the user json byte : ", buser)
	fmt.Println("This is the user json string : ", string(buser))

	//////////////////////////
	//unmarshal the users:
	s2 := `[{"Name":"Steven","Age":30,"Sayings":["You are awesome","You are great"]},{"Name":"Kedu","Age":40,"Sayings":["This is him","He is an awesome teacher"]}]`
	//first convert is to slice of byte:
	bs2 := []byte(s2)
	fmt.Println("this is the json slice of byte: ", bs2)
	fmt.Printf("%T\n", bs2)
	//create a slice user:
	var users2 []userNew
	erruser = json.Unmarshal(bs2, &users2)
	if erruser != nil {
		fmt.Println("Error occured: ", erruser)
	}
	fmt.Println("This is the normal user: ", users2)
	fmt.Println()

	for i, v := range users2 {
		fmt.Println("For the user: ", i)
		fmt.Println("\t", v.Name)
		for _, val := range v.Sayings {
			fmt.Println("\t\t", val)
		}
	}

	//////////////////////////
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
