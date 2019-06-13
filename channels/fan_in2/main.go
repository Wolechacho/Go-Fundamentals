package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//c is a receive only channel
	c := fanIn(boring("Joe"), boring("Ann"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You are both boring; i'm leaving")
}

//this is a receive only channel
func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

//Fanin
//fanin have two receive channels
//ie put the values of two channels into one channel
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			//take the value from input1 and put in the the channel c
			c <- <-input1
		}
	}()
	go func() {
		for {
			//take the value from input2 and put in the the channel c
			c <- <-input2
		}
	}()
	return c
}
