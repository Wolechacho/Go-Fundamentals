package main

import "fmt"

func main() {
	fmt.Println("Welcome to testing")
	fmt.Println("2 + 3 = ", mySum(2, 3))
	fmt.Println("4 + 7 = ", mySum(4, 7))

}

//a pointer of type T from the package testing
// func TestAverage(t *testing.T){
// }

//a variadic parameter is x
func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
