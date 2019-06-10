package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name    string
	Age     int
	Sayings []string
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
	fmt.Println(users)

	//turning the users data to json
	err := json.NewEncoder(os.Stdout).Encode(users)

	if err != nil {
		fmt.Println("We have an error: ", err)
	}
}
