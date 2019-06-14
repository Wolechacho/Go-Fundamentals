package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {

	//launch 10 go routines and put 10 numbers in each
	//observe that we didnt use close here
	// c := make(chan int)
	// for j := 0; j < 10; j++ {
	// 	go func() {
	// 		for i := 0; i < 10; i++ {
	// 			c <- i
	// 		}
	// 	}()
	// }
	// for k := 0; k < 100; k++ {
	// 	fmt.Println(k, <-c)
	// }
	// fmt.Println("about to exit")

	//Solution 2
	// 	x := 10
	// 	y := 10
	// 	c := gen(x, y)

	// 	for i := 0; i < x*y; i++ {
	// 		fmt.Println(i, <-c)
	// 	}
	// 	fmt.Println("Routines", runtime.NumGoroutine())
	// }
	// func gen(x, y int) <-chan int {
	// 	c := make(chan int)

	// 	for i := 0; i < x; i++ {
	// 		go func() {
	// 			for j := 0; i < y; j++ {
	// 				c <- j
	// 			}
	// 		}()
	// 		fmt.Println("Routines", runtime.NumGoroutine())
	// 	}
	// 	return c

	c := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				c <- rand.Intn(1000)
			}
			wg.Done()
		}()
	}

	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()
	wg.Wait()
	close(c)
	fmt.Println("exiting")
}
