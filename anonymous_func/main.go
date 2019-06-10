package main

import "fmt"

func main() {

	func() {
		fmt.Println("Anonymous without argument")
	}()
	func(x int) {
		fmt.Println("Anonymous with x:", x)
	}(32)
}
