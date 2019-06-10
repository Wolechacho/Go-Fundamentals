package main

import (
	"fmt"
	"sort"
)

type user struct {
	Name    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []user

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	user1 := user{
		Name: "Steven",
		Age:  30,
		Sayings: []string{
			"You are awesome",
			"Great are great",
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

	user3 := user{
		Name: "Eke",
		Age:  500,
		Sayings: []string{
			"Joyous man",
			"Incredible",
		},
	}

	users := []user{user1, user2, user3}
	fmt.Println(users)
	sort.Sort(ByAge(users))
	fmt.Println(users)
	sort.Sort(ByName(users))
	fmt.Println(users)

	for i, v := range users {
		fmt.Println("For user: ", i)
		fmt.Println("\t", v.Name, v.Age)
		sort.Strings(v.Sayings)
		for _, val := range v.Sayings {
			fmt.Println("\t\t", val)
		}
	}
}
