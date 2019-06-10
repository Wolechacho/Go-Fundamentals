package main

import "fmt"

func main() {

	f := func() {
		fmt.Println("this is the func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("this is the func expression", x)
	}
	g(40)
}
