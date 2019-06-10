package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Before cpu", runtime.NumCPU())
	fmt.Println("Before gs", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hello from 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from 2")
		wg.Done()

	}()

	fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("about to exit")
	fmt.Println("After cpu", runtime.NumCPU())
	fmt.Println("After gs", runtime.NumGoroutine())
}
