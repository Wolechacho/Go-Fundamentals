package main

import "fmt"

var y = 69

var z = `peter said that : "This is the one"`

var a int

type hotdog int

var b hotdog

func main() {

	a = 43
	b = 44
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = hotdog(a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	// fmt.Println("Hello")
	// for i := 0; i < 100; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }
	// x := 42
	// fmt.Println(x)
	// fmt.Println(y)

	// fmt.Printf("%T\n", y)
	// fmt.Printf("%b\n", y)
	// fmt.Printf("%x\n", y)
	// fmt.Printf("%#x\n", y)

	// fmt.Println(z)
	// fmt.Printf("%T\n", z)

}
