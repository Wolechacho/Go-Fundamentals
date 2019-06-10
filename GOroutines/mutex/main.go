package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Routines : ", runtime.NumGoroutine())

	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	//adding mutex
	//Which is used to lock down the code to prevent race condition
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Routines : ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("count", counter)
}
