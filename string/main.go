package main

import "fmt"

func main() {
	s := "Message me"
	fmt.Println(s)
	a := `"Awesome
	Friend"`
	fmt.Println(a)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
	fmt.Println(" ")

	for i, v := range s {
		// fmt.Println(i, v)
		fmt.Printf("at index position of %d, we have %#x\n", i, v)
	}
}
