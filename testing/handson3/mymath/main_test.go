package mymath

import (
	"fmt"
	"testing"
)

func ExampleCenterAvg() {
	fmt.Println(CenterAvg([]int{2, 4, 5, 9, 10}))
}

func TestCenterAvg(t *testing.T) {
	//type "avg" underline type of struct
	type avg struct {
		data []int
		ans  float64
	}

	sampleData := []avg{
		avg{[]int{20, 30, 40, 50}, 35},
		avg{[]int{900, 20, 70, 50}, 60},
		avg{[]int{200, 5, 40, 7, 9}, 18.666666666666668},
	}
	for _, v := range sampleData {
		x := CenterAvg(v.data)
		if x != v.ans {
			t.Error("Got", x, "want", v.ans)
		}
	}
}

func BenchmarkCenterAvg(b *testing.B) {
	xi := []int{20, 30, 54, 90}
	for i := 0; i < b.N; i++ {
		CenterAvg(xi)
	}
}
