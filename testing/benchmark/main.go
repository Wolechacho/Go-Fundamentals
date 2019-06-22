package main

import (
	"fmt"
	"strings"

	"./saying"
	"./str"
)

const s = "this is the string"

func main() {
	fmt.Println(saying.Greet("Sam"))
	//from the strings package, we call Split
	//we get all stufs separated by space and prints them in a new line
	xs := strings.Split(s, " ")

	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", str.Cat(xs))
	fmt.Printf("\n%s\n", str.Join(xs))

}
