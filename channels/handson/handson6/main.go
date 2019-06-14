package main

import "fmt"

func main() {

	c := make(chan int)

	//this is one goroutine, that loop we have a loop to insert 100 numbers
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		//without the close, we will hava a deadlock issue
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("about to exit")
}
