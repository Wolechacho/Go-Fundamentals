package main

import "fmt"

var g = func() {
	fmt.Println("THis is the g print")
}

func main() {

	s := func() {
		fmt.Println("This is the function assigned")
	}
	s()

	fmt.Printf("%T\n", s)

	g()

	theFunc := mike()

	fmt.Println("this is the  print: ", theFunc())

	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}

	x := foo(g, []int{1, 2, 3, 4, 5, 6})
	fmt.Println("this is the x: ", x)
}

func mike() func() int {
	return func() int {
		return 43
	}
}

func foo(f func(x []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}
