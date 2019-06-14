package main

import "fmt"

type customError struct {
	info string
}

//Any value of type customError implements the error interface
func (ce customError) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func main() {
	c1 := customError{
		info: "Need more Coffee",
	}
	foo(c1)
}

func foo(e error) {
	//e.info will not work
	//we can achieve that using assertion(not conversion)
	//checking an interfaces
	// fmt.Println("Foo ran - ", e)
	fmt.Println("Foo ran - ", e.(customError).info)

}
