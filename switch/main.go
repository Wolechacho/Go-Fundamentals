package main

import "fmt"

func main() {
	//since we dont have value on the switch, it evaluates based on bool
	switch {
	case false:
		fmt.Println("this wont print")
	case (2 == 2):
		fmt.Println("This is correct")
		fallthrough
	case (3 == 3):
		fmt.Println("This is also correct")
		fallthrough
	case (3 == 4):
		fmt.Println("This is not correct")
		fallthrough
	default:
		fmt.Println("this is the default")
	}
	//Now lets have a value for the switch statement
	n := "Bond"
	switch n {
	case "MoneyPenny", "Bond":
		fmt.Println("this money also James Bond")
	case "M":
		fmt.Println("for M")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is the default")
	}
}
