package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
	}
	defer f2.Close()
	fmt.Println("Check the log.txt file in the directory")

	defer foo()

	//Using Fatalln
	// _, err = os.Open("no-one.txt")
	// if err != nil {
	// 	//the program exit with status 1, defer functions wont run
	// 	log.Fatalln(err)
	// }

	//Using Panic
	// if err != nil {
	// 	//the program exit with status 1, defer functions wont run
	// 	log.Panicln(err)
	// }
}

func foo() {
	fmt.Println("Hello")
}
