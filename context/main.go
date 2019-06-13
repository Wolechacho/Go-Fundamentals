package main

import (
	"context"
	"fmt"
)

func main() {

	ctx := context.Background()
	fmt.Println("Context: \t", ctx)
	fmt.Println("Context err: \t", ctx.Err())
	fmt.Printf("Context type: \t%T\n", ctx)
	fmt.Println("------------")

	ctx, _ = context.WithCancel(ctx)
	fmt.Println("Context: \t", ctx)
	fmt.Println("Context err: \t", ctx.Err())
	fmt.Printf("Context type: \t%T\n", ctx)
	fmt.Println("------------")

	ctxa, cancel := context.WithCancel(ctx)
	fmt.Println("Context: \t", ctxa)
	fmt.Println("Context err: \t", ctxa.Err())
	fmt.Printf("Context type: \t%T\n", ctxa)
	fmt.Println("cancel: \t\t", cancel)
	fmt.Printf("cancel type: \t%T\n", cancel)

	fmt.Println("------------")

	cancel()

	fmt.Println("Context: \t", ctxa)
	fmt.Println("Context err: \t", ctxa.Err())
	fmt.Printf("Context type: \t%T\n", ctxa)
	fmt.Println("------------")

}
