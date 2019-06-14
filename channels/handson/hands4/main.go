package main

import "fmt"

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()
	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		//grab values from the c channel
		case v := <-c:
			fmt.Println(v)
			//if we have values in q, return(we are done)
		case <-q:
			// fmt.Println("this is for q")
			//if this return is not here, we wil have a bunch of zeros
			return
		}
	}
}
