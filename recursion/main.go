package main

import "fmt"

func main() {

	fmt.Println(factorial(4))
	fmt.Println(loopFac(5))

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFac(n int) int {
	total := 1
	// for i := 1; i <= n; i++ {
	// 	total *= i
	// }
	//or
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
