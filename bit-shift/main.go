package main

import "fmt"

const (
	_ = iota
	// kb = 1024
	// mb = 1024 * kb
	// gb = 1024 * mb
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	// x := 5
	// fmt.Printf("%d\t\t%b\n", x, x)
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}
