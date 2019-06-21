package main

import "testing"

func TestMySum(t *testing.T) {

	//TABLE OF TEST
	type test struct {
		data   []int
		answer int
	}
	//this is the a slice of values of type test
	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{3, 4, 5}, 12},
		test{[]int{1, 1}, 2},
		test{[]int{-1, 0, 1}, 0},
	}
	for _, v := range tests {
		//"Unfuring" the variadic parameter
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}

	//this is for the normal test
	// x := mySum(2, 3)
	// if x != 5 {
	// 	t.Error("Expected", 5, "Got", x)
	// }

}

// this is the normal test
func TestMySub(t *testing.T) {
	x := mySub(7, 4)
	if x != 3 {
		t.Error("Expected", 3, "Got", x)
	}
}
