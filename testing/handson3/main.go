package main

import (
	"fmt"

	"./mymath"
)

func main() {
	xxi := gen()
	for _, v := range xxi {
		fmt.Println("this is the ans: ", mymath.CenterAvg(v))
	}
}

//We are returning a slice of slice of int
func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	//[][]int means a slice of a slice of int
	e := [][]int{a, b, c, d}

	return e
}
