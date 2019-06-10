package main

import "fmt"

func main() {
	// for i := 0; i <= 4; i++ {
	// 	fmt.Printf("The outer loop: %d\n", i)
	// 	for j := 0; j < 10; j++ {
	// 		fmt.Printf("\t\t The inner loop: %d\n", j)
	// 	}
	// }

	// x := 1
	// for x < 10 {
	// 	fmt.Println(x)
	// 	x++
	// }
	// fmt.Println("done.")

	// for {
	// 	x++
	// 	if x > 9 {
	// 		fmt.Println("we are done")
	// 		break
	// 	}
	// 	//e.g, 3%2 != 0, so, skip that value 3 and every similar
	// 	if x%2 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println(x)
	// }

	// for i := 65; i <= 90; i++ {
	// 	fmt.Printf("%d\n", i)
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Printf("\t%#U\n", i)
	// 	}
	// }

	// bd := 1990
	// for bd <= 2019 {
	// 	fmt.Println(bd)
	// 	bd++
	// }

	// for {
	// 	if bd > 2019 {
	// 		break
	// 	}
	// 	fmt.Println(bd)
	// 	bd++
	// }

	for i := 10; i < 100; i++ {
		//remove cases when the remainder is zero
		if i%4 != 0 {
			fmt.Printf("When %v is divided by 4, the remainder is %v\n", i, i%4)
		}
	}
}
