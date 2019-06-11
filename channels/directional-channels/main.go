package main

import "fmt"

func main() {

	// c := make(chan int, 2)

	//this is a receive only channel
	//we cant semd to this channel
	//we are receiving from a channel of int
	// c := make(<-chan int, 2)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println("-----------")
	// fmt.Printf("%T\n", c)

	//////////////////////////////////////
	// c := make(chan int)
	// cr := make(<-chan int) //receive
	// cs := make(chan<- int) //send
	// fmt.Println("-----------")
	// fmt.Printf("c\t%T\n", c)
	// fmt.Printf("cr\t%T\n", cr)
	// fmt.Printf("cs\t%T\n", cs)

	/////////////////////////////////////////
	//THIS WILL NOT WORK, BECAUSE, THIS CHANNEL WILL ONLY SEND AND NOT RECEIVE
	// cs := make(chan<- int) //send
	// go func() {
	// 	//send 42 to the channnel
	// 	cs <- 42
	// }()
	// fmt.Println(<-cs) //when we try to receive => error
	// fmt.Printf("-----------\n")
	// fmt.Printf("cs\t%T\n", cs)

	/////////////////////////////////////////
	//THIS WILL NOT WORK, BECAUSE, THIS CHANNEL WILL ONLY RECEIVE AND NOT SEND
	// cr := make(<-chan int) //receive
	// go func() {
	// 	//send 42 to the channnel
	// 	cr <- 42 //when we try to send => error
	// }()
	// fmt.Println(<-cr)
	// fmt.Printf("-----------\n")
	// fmt.Printf("cr\t%T\n", cr)

	/////////////////////////////////////
	// Also, specific to general doesnt work
	// c = cs
	// c = cr

	// Also, general to specific will work
	// cs = c
	// cr = c

	//////////////////////////////////////
	c := make(chan int)
	// cr := make(<-chan int) //receive
	// cs := make(chan<- int) //send

	//general to specific converts
	fmt.Println("-----------")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))

	//specific to general wont convert
	// fmt.Println("-----------")
	// fmt.Printf("c\t%T\n", (chan int)(cr))
	// fmt.Printf("c\t%T\n", (chan int)(cs))
}
