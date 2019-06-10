package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//since we just want integers not slices, thats why we need to "..."
	s := sum(ii...)
	fmt.Println("All numbers ", s)

	//note: sum is a function
	s2 := even(sum, ii...)
	fmt.Println("Even numbers ", s2)

	s3 := odd(sum, ii...)
	fmt.Println("Odd numbers ", s3)

}

//variadic int
func sum(xi ...int) int {
	fmt.Printf("%T\n", xi) //this will give us slice of int
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

//Callback: we pass a function as an argument to another function

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
