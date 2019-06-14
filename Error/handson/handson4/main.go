package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat string
	lng string
	err error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.lng, se.err)
}
func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Print(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// e := errors.New("more coffee here")
		e := fmt.Errorf("more coffee here - value was %v", f)
		return 0, sqrtError{"20.34 N", "90.54 W", e}
	}
	return 43, nil
}
