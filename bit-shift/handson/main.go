package main

import "fmt"

const (
	a = 2017 + iota
	b = 2017 + iota
	c = 2017 + iota
	d = 2017 + iota
)

func main() {
	x := 42
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)

	y := x << 1
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
