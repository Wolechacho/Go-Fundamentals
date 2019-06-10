package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)

	//shortest way
	fmt.Println(bar()())

	x := bar()
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)
}

func foo() string {
	s := "Hello Sir"
	return s
}

//note, the return type here is func() int
func bar() func() int {
	return func() int {
		return 432
	}
}
