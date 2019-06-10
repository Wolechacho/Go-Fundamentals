package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

func main() {

	xi := []int{10, 19, 59, 2, 3, 5, 6, 7, 8}
	xs := []string{"Zara", "Magi", "James", "John", "Mike", "Ojo"}

	fmt.Println(xi)
	fmt.Println(xs)
	sort.Ints(xi)
	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)

	p1 := person{"James", 23}
	p2 := person{"Ade", 53}
	p3 := person{"Mark", 233}
	p4 := person{"Alo", 293}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)

}
