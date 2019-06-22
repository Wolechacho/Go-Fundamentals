package mymath

import "sort"

func CenterAvg(xi []int) float64 {
	sort.Ints(xi)
	//we slice the slice of int
	//drop the lowest value, drop the highest value
	//slicing from 1 to len of xi minus 1
	//the reason we use minus 1 to mean the last element of the array is because, the index of the array started from zero

	//The code below states: get from the first to, not including the last one
	//we leave out the first and the last guys
	//if we have {1,2,3,4,5}, drop 1 and 5, we have 2,3,4. Add them and divide by their count
	xi = xi[1:(len(xi) - 1)]

	n := 0
	for _, v := range xi {
		n += v
	}

	f := float64(n) / float64(len(xi))
	return f
}
