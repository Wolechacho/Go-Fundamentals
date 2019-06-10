package main

import "fmt"

func main() {

	//func (r receiver) identifier(parameters) (return(s)) { ... }
	foo()
	bar("James")
	s1 := woo("Wonderful")
	fmt.Println(s1)

	//note, we are assigning the returns
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)

	xi := []int{2, 3, 4, 5, 6, 7}
	// summing := sumFunc(2, 3, 4, 5, 6, 7)
	//since xi is a slice of int and the function "sumFunc" original type is "int" we need to add the "..."
	summing := sumFunc("james", xi...)
	fmt.Println("The total is: ", summing)

	fmt.Println()

	// when we pass no values
	summing2 := sumFunc("Mark")
	fmt.Println("The total is: ", summing2)

	fmt.Println()

	defer aki()
	popo()

	mark, ali := hello()
	fmt.Println("this is the integer: ", mark)
	fmt.Println("this is the string: ", ali)

	ii1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	n := test1(ii1...)
	fmt.Println("THis is the summation: ", n)

	//this is the second summation:
	m := test2(ii1)
	fmt.Println("THis is the second summation: ", m)

}

func foo() {
	fmt.Println("This is the print")
}

func bar(s string) {
	fmt.Println("Hello, ", s)
}

func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true
	return a, b
}

/////////////////////////////////////////////
//Variadic parameters
//x has unlimited number of ints
// ...int is vaiadic. it means zero or more of type int
func sumFunc(s string, x ...int) int {
	//the output will be a slice of int
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0
	for _, v := range x {
		sum += v
		// fmt.Println("for item in index position ", i, "we are now adding ", v, "to the total which is now ", sum)
	}

	// fmt.Println("The total is, ", sum)
	return sum
}

func aki() {
	fmt.Println("Aki")
}

func popo() {
	fmt.Println("Popo")

}

//return a string and an integer:

func hello() (int, string) {
	return 123, "My big brother"
}

func test1(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func test2(xi []int) int {
	total := 0

	for _, v := range xi {
		total += v
	}
	return total
}
