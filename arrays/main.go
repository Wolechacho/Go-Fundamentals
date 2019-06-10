package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)

	y := [5]int{2, 3, 5, 6, 7}

	for i, v := range y {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", y)
}
