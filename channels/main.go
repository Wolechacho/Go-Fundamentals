package main

import "fmt"

func main() {
	//this is called making a buffer channel
	c := make(chan int, 2)
	c <- 43
	c <- 46
	// go func() {
	// 	c <- 43
	// 	c <- 45
	// }()
	// take off the value from the channel
	fmt.Println(<-c)
	fmt.Println(<-c)

	// c := make(chan int)
	// go func() {
	// 	c <- 43
	// 	c <- 45
	// }()
	// // take off the value from the channel
	// fmt.Println(<-c)

}
