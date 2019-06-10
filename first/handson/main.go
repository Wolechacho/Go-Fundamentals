package main

import "fmt"

var x = 42
var y string = "James Bond"
var z bool

type mongo int

var a mongo
var b int

func main() {
	//This will print zero values
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)

	fmt.Println(s)

	fmt.Println(a)
	a = 50
	fmt.Printf("%T\n", a)
	fmt.Println(a)

	b = int(a)
	fmt.Println(b)

}
