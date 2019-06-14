package main

import (
	"fmt"
	"log"
)

type negativeMathError struct {
	lat  string
	long string
	err  error
}

func (n negativeMathError) Error() string {
	return fmt.Sprintf("a norgate math error  occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate with math redux: square root of negative number: %v", f)
		return 0, fmt.Errorf("norgate math again: square root of negative number: %v", nme)
	}
	return 42, nil
}
