package main

import "fmt"

func main() {
	c := make(chan int)
	//this anonymous func prevent deadlock,
	go func() {
		//here, we are sending 43 to c
		c <- 43
	}()
	//remember the code below means, we are receiving from the c channel
	fmt.Println(<-c)

	//using buffer channel
	d := make(chan int, 1)
	//here, we are sending 43 to d
	d <- 43
	//remember the code below means, we are receiving from the d channel
	fmt.Println(<-d)
}
